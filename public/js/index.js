"use strict";

$(document).ready(function () {

  $('select').selectpicker();

  let isBlank = function (string) {
    return string == null || string === "";
  };
  
  $('.btn_clean').on('click', function () {
    location.href = "/"
  });

  $('.btn_print').on('click', function () {
    let thisbtn = $('.btn_print');
    let testHtml = $('div.testHtml');
    
    $.ajax({
      type: "POST",
      url: "/happy",
      data: {},
      dataType: "html",
    })
    .done(async function (data) {
      // Add data in Modal body
      //thisbtn.remove();
      testHtml.html(data);
    })
    .fail(function (jqXHR, textStatus, errorThrown) {
      console.log('서버오류: '+ textStatus);
      console.log('서버오류: '+ errorThrown);
    })
  });

  // $('.btn_print').on('click', function () {
  //   //$("div.testHtml").html(`<h2>test</h2>`);
  //   $("div.testHtml").append(`<h2>hello test</h2>`);
  // });

  
  // $('.btn_print').on('clack', function(event) {
  //   event.preventDefault();
  //   alert("hi")
  //   $(".test").attr("class");
  //   $(".test").append(`<h2>hello test</h2>`);
  // });
  
  /**
   * 대학 List의 대학 선택시
   * 대학에 대한 정보 보여줌
   */
  $(document).on("click", "tr", function (e) {
  // $('td.univId').on('click', function (e) {   
    e.preventDefault();
    var id = $(this).children('.univId').attr('id');
    
    if (id != "") {
      $.ajax({
        type: "POST",
        url: "/univDetail",
        data: { id },
        dataType: "html",
      })
      .done(async function (response) {
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
    } else 
    if (isBlank(fAvgValue)) {
      alert("내신 평균을 입력하세요.");
      $('input#fAvgValue').focus();
    } else if (isBlank(lAvgValue)) {
      alert("내신 평균을 입력하세요.");
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

  // let username;
  // while (isBlank(username)) {
  //   username = prompt("What's your name?");
  //   if (!isBlank(username)) {
  //     $('#user-name').html('<b>' + username + '</b>');
  //   }
  // }

});