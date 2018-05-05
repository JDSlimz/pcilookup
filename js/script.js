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
