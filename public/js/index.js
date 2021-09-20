"use strict";

$(document).ready(function () {

  $('select').selectpicker();  

  let isBlank = function (string) {
    return string == null || string === "";
  };
  
  $('.btn_clean').on('click', function () {
    location.href = "/"
  });

  // Print all univ Info only checked checkbox
  $('.btn_print').on('click', function (e) {
    e.preventDefault();

    let chk_val = [];
    $('input.chkList:checked').each(function(index) {
      chk_val.push($(this).val());
    });
    //alert(chk_val);
    if (chk_val=="") {
      alert("지원대학을 선택하세요.");
    } else {

      $.ajax({
        type: "POST",
        url: "/printUnivInfo",
        data: { chk_val },
        dataType: "html",
      })
      .done(async function (data) {
        $(location).attr("href", "/excel");
      })
      .fail(function (jqXHR, textStatus, errorThrown) {
        console.log('서버오류: '+ textStatus);
        console.log('서버오류: '+ errorThrown);
      })
    }
  });

  
  /**
   * 대학 List의 대학 선택시
   * 대학에 대한 정보 보여줌
   */
  $(document).on("click", 
  "td.id, td.univId, td.univForm, td.dept, td.line, td.volume, td.name, td.1st, td.1stVol, td.final, td.sulimit", 
  function (e) {
    //$(document).on("click", "tr", function (e) {
    // $('td.univId').on('click', function (e) {   
    e.preventDefault();
    
    let tr = $(this).closest("tr");
    let id = tr.children('.univId').attr('id');
    // let id = $(this).children('.univId').attr('id');
    
    // 자신의 checkbox check여부 확인
    let idVal = "false";
    if($(`input#chkList_${id}`).is(':checked')){ 
      idVal = "true";
    } else {
      idVal = "false";
    }

    if (id != "") {
      $.ajax({
        type: "POST",
        url: "/univDetail",
        data: { id },
        dataType: "html",
      })
      .done(async function (response) {
        // 모달창 checkbox 반영
        if(idVal == "true"){
          $('input.subChkList').prop('checked', true);
        } else {
          $('input.subChkList').prop('checked', false);
        }

        // Add response in Modal body
        $('.modal-body').html(response); 
    
        // Display Modal
        $('#univModal').modal('show');
        
      })
      .fail(function (jqXHR, textStatus, errorThrown) {
        console.log('서버오류: '+ textStatus);
        console.log('서버오류: '+ errorThrown);
      })
    }
  });
  // !$('input[type="checkbox"]:checked').each(function(index) {
  //   callId = $('.univId').attr('id');
  //   if (id == callId) {
  //     $('input:checkbox').prop("checked");


  // 전형, 계열, 전공 구분 
  $('.btn_serchDept').on('click', function (e) {
    e.preventDefault();
    let getApplyForm = $('.selectpicker.applyForm').val(); // 전형구분 
    let getApplyLine = $('.selectpicker.applyLine').val(); // 계열구분
    let getApplySubject = $('.selectpicker.applySubject').val(); // 전공구분
    
    //alert(getApplyForm+":"+getApplyLine+":"+getApplySubject);
    // console.log(getApplyForm);
    // console.log(getApplyLine);
    // console.log(getApplySubject);

    if (getApplyForm =="") {
      alert("전형구분를 선택하세요.");
    } else if (getApplyLine=="") {
      alert("계열구분를 선택하세요.");
    } else if (getApplySubject=="") {
      alert("전공구분를 선택하세요.");
    } else {

      $.ajax({
        type: "POST",
        url: "/indexP",
        data: { getApplyForm, getApplyLine, getApplySubject},
        dataType: "html",
        async: false,
      })
      .done(function (data) {
        // Add data in 모집단위
        $("select.applyDept option").remove();
        $("select.applyDept").append(data);
        $(".selectpicker").selectpicker('refresh');

        if($("button.btn_serchUniv").prop("disabled")){
          $("button.btn_serchUniv").prop("disabled", false);
        }
        $("div.univListbydept").html("");
        $('input#myValue').val("");
        $('input#fAvgValue').val("");
        $('input#lAvgValue').val("");
        $("section.inValSrh").prop("hidden", true);
        $("input#img").prop("hidden", false);
        $('.btn_print').prop("hidden", true);
      })
      .fail(function (jqXHR, textStatus, errorThrown) {
        console.log('서버오류: '+ textStatus);
        console.log('서버오류: '+ errorThrown);
      })
    }
  });

  // 모집단위 결과
  $('.btn_serchUniv').on('click', function (e) {
    e.preventDefault();
    let selectedItem = $('.selectpicker.applyDept').val();
    //alert(selectedItem);
    if (isBlank(selectedItem)) {
      alert("모집단위를 선택하세요.");
    } else {

      $.ajax({
        type: "POST",
        url: "/indexPuniv",
        data: { selectedItem },
        dataType: "html",
      })
      .done(async function (data) {
        // Add data in 대학목록
        //$('div.univListChart').html(`<iframe width="1200" height="550" frameborder="0" style="border:0" src="/page" allowfullscreen></iframe>`);        
        $("input#img").prop("hidden", true);
        $('div.univListbydept').html(data);
        $("section.inValSrh").prop("hidden", false);
        $('.btn_print').prop("hidden", false);

      })
      .fail(function (jqXHR, textStatus, errorThrown) {
        console.log('서버오류: '+ textStatus);
        console.log('서버오류: '+ errorThrown);
      })
    }
  });

  // 내점수
  $('.btn_serch_myValue').on('click', function (e) {
    e.preventDefault();
    
    let selectedItem = $('.selectpicker.applyDept').val();
    //alert(selectedItem);

    let myValue = $('input#myValue').val();
    //alert(myValue);

    if (isBlank(selectedItem)) {
      alert("모집단위를 선택하세요.");
    } else if (isBlank(myValue)) {
      alert("내 점수를 입력하세요.");
    } else if (myValue<0 || myValue>9) {
      alert("내 점수의 범위는 0~9 사이입니다.");
      $('input#myValue').focus();
    } else {

      $.ajax({
        type: "POST",
        url: "/indexPuniv",
        data: { selectedItem, myValue },
        dataType: "html",
      })
      .done(async function (data) {
        // Add data in 대학목록
        //$('div.univListChart').html(`<iframe width="1200" height="550" frameborder="0" style="border:0" src="/page" allowfullscreen></iframe>`);        
        $('div.univListbydept').html(data);

      })
      .fail(function (jqXHR, textStatus, errorThrown) {
        console.log('서버오류: '+ textStatus);
        console.log('서버오류: '+ errorThrown);
      })
    }
  });

  // 2021 내신평균
  $('.btn_serch_avg').on('click', function (e) {
    e.preventDefault();
    
    let selectedItem = $('.selectpicker.applyDept').val();
    //alert(selectedItem);

    let fAvgValue = $('input#fAvgValue').val();
    let lAvgValue = $('input#lAvgValue').val();
    //alert(`${fAvgValue}:${lAvgValue}`);
    
    if (isBlank(selectedItem)) {
      alert("모집단위를 선택하세요.");
    } else if (isBlank(fAvgValue)) {
      alert("내신 평균을 입력하세요.");
      $('input#fAvgValue').focus();
    }  else if (fAvgValue<0 || fAvgValue>9) {
      alert("내신 평균의 범위는 0~9 사이입니다.");
      $('input#fAvgValue').focus();
    } else if (isBlank(lAvgValue)) {
      alert("내신 평균을 입력하세요.");
      $('input#lAvgValue').focus();
    } else if (lAvgValue<0 || lAvgValue>9) {
      alert("내신 평균의 범위는 0~9 사이입니다.");
      $('input#lAvgValue').focus();
    } else {

      $.ajax({
        type: "POST",
        url: "/indexPuniv",
        data: { selectedItem, fAvgValue, lAvgValue},
        dataType: "html",
      })
      .done(async function (data) {
        // Add data in 대학목록
        //$('div.univListChart').html(`<iframe width="1200" height="550" frameborder="0" style="border:0" src="/page" allowfullscreen></iframe>`);        
        $('div.univListbydept').html(data);

      })
      .fail(function (jqXHR, textStatus, errorThrown) {
        console.log('서버오류: '+ textStatus);
        console.log('서버오류: '+ errorThrown);
      })
    }
  });
  
  // 내점수 입력란 값 삭제
  $('button.btn_clean_myValue').on('click', function () {
    $('input#myValue').val("");
  });

  // 2021 내신 평균 입력란 값 삭제
  $('button.btn_clean_avgValue').on('click', function () {
    $('input#fAvgValue').val("");
    $('input#lAvgValue').val("");
  });

  // checkbox 대학학과 리스트
  $(document).on("click", "button#univInfoCheckBox", function () {    
    $('input:checkbox').prop("checked", !$('input:checkbox').prop("checked"));
  });

 // 대학 상세 화면 에서 체크박스 이벤트
 // 모달창 체크여부에 따라 리스트의 체크 여부 반영됨
  $('input.subChkList').on('change', function () {
    
    let id = $('.modal_univId').attr('id');
    // alert(`chkList_${id}`);
    if($('input.subChkList').is(':checked')){
      $(`input#chkList_${id}`).prop('checked', true);
    } else {
      $(`input#chkList_${id}`).prop('checked', false);
    }
  });

});