$(document).ready(function(){

	var ven = getParameterByName("ven");
	var dev = getParameterByName("dev");

	$.ajax({
		url: "https://api.pcilookup.com",
		data: {action:'search', vendor: ven, device: dev},
		dataType: 'json',
		success: function(data) {
			var table = $('#results').DataTable( {
				"paging": false,
				"columns": [
				    { "className": "hex-id" },
				    { "className": "description" }
				],
			} );

			$.each(data, function(i, item) {
			    $('#manufacturer').html(item.Description);
			    if (typeof item.Devices !== 'undefined') {
				    var devs = item.Devices.length;
				    if(devs > 0){
				    	$.each(item.Devices, function(i, item){
				    		var row = table.row.add([
			    				item.ID,
			    				item.Desc
			    			]);
				    		if('Sub' in item){
				    			$.each(item.Sub, function(i, subItem){
				    				var rowData = row.child();

				    				if(rowData == null){
				    					rowData += subItem.ID + ':' + subItem.Desc + '<br>';
				    				} else {
				    					rowData = subItem.ID + ':' + subItem.Desc + '<br>';
				    				}
				    				row.child(rowData);
				    			});
				    		}
				    	});
				    }
				}
			});
			table.draw();

			$('#results tbody').on('click', 'td.details-control', function () {
		        var tr = $(this).closest('tr');
		        var row = table.row( tr );
		 
		        if ( row.child.isShown() ) {
		            // This row is already open - close it
		            row.child.hide();
		            tr.removeClass('shown');
		        }
		        else {
		            // Open this row
		            row.child().show();
		            tr.addClass('shown');
		        }
		    } );
		    $('#processingSpinner').hide();
		},
		error: function(xhr, status, error) {
			var err = eval("(" + xhr.responseText + ")");
			console.log(xhr);
			$('#processingSpinner').hide();
		}
	});
});

function getParameterByName(name, url) {
	if (!url) url = window.location.href;
	name = name.replace(/[\[\]]/g, "\\$&");
	var regex = new RegExp("[?&]" + name + "(=([^&#]*)|&|#|$)"), results = regex.exec(url);
	if (!results) return null;
	if (!results[2]) return '';
	return decodeURIComponent(results[2].replace(/\+/g, " "));
}
