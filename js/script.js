$(document).ready(function(){

	var ven = getParameterByName("ven");
	var dev = getParameterByName("dev");

	$('#results').DataTable( {
		ajax: "/api.php?action=search&vendor=" + ven + "&device=" + dev,
		columns: [
			{ data: 'venDesc' },
			{ data: 'venID' },
			{ data: 'desc' },
			{ data: 'id' }
		],
		"sAjaxDataProp": "",
		responsive: true
	} );
});

function getParameterByName(name, url) {
	if (!url) url = window.location.href;
	name = name.replace(/[\[\]]/g, "\\$&");
	var regex = new RegExp("[?&]" + name + "(=([^&#]*)|&|#|$)"), results = regex.exec(url);
	if (!results) return null;
	if (!results[2]) return '';
	return decodeURIComponent(results[2].replace(/\+/g, " "));
}

function submitSurvey(){
	var ableToFind = $('input[name=ableToFind]:checked').val();
	var deviceInfo = $('#deviceInfo').val();
	var overallRating = $('input[name=overallRating]:checked').val();
	var comments = $('#comments').val();
	
	if(ableToFind == null){
		$('#ableToFindContainer').addClass('error');
	} else {
		$('#ableToFindContainer').removeClass('error');
	}
	if(overallRating == null){
		$('#overallRatingContainer').addClass('error');
	} else {
		$('#overallRatingContainer').removeClass('error');
	}
	$.get("/api.php", $( "#surveyForm" ).serialize(), function(){
		$('#surveyModal').modal('hide');
		$('#manufacturer').before('<span id="message" style="display:block; width:100%; text-align:center; background-color:green; color:white">Survey submitted! Thank you for your feedback!</span>');
		setTimeout(fade_out, 5000);
	} );
}

function fade_out() {
	$("#message").fadeOut().empty();
}