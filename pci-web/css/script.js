$(document).ready(function(){

	var ven = getParameterByName("ven");
	var dev = getParameterByName("dev");

	$('#results').DataTable( {
		ajax: "https://dev.pcilookup.com?action=search&vendor=" + ven + "&device=" + dev,
		columns: [
			{ data: 'Vendor Description' },
			{ data: 'Vendor ID' },
			{ data: 'Desc' },
			{ data: 'ID' }
		],
		"sAjaxDataProp": "",
		responsive: true
	} );

	/*$.ajax({
		url: "https://dev.pcilookup.com",
		data: {action:'search', vendor: ven, device: dev},
		dataType: 'json',
		success: function(data) {
			var table = $('#results').DataTable( {
				"paging": false,
				"columns": [
				    { "className": "hex-id" },
				    { "className": "description" }
				],
			});

			var vendors = {};
			var devices = {};

			$.each(data, function(i, item) {

				var vendorID = item.VendorID;
				var vendorDesc = item.VendorDescrip;
				var devID = item.ID;
				var devDesc = item.Desc;

				var vendor = {vendorID:vendorDesc};

				if(vendors.indexOf(vendor) == -1){
					vendors.push(vendor);
				}

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
	});*/
});

function getParameterByName(name, url) {
	if (!url) url = window.location.href;
	name = name.replace(/[\[\]]/g, "\\$&");
	var regex = new RegExp("[?&]" + name + "(=([^&#]*)|&|#|$)"), results = regex.exec(url);
	if (!results) return null;
	if (!results[2]) return '';
	return decodeURIComponent(results[2].replace(/\+/g, " "));
}
