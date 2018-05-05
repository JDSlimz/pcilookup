<?php

?>
<form id="search" class="form-horizontal">

	<!--Vendor Input-->
	<div class="form-group">
		<label class="control-label" for="ven">Vendor</label>  
		<div class="">
			<input id="ven" name="ven" placeholder="e.g. 10de or Nvidia" class="form-control input-lg" type="text">
		</div>
	</div>

	<!--Device Input-->
	<div class="form-group">
		<label class="control-label" for="dev">Device</label>  
		<div class="">
			<input id="dev" name="dev" placeholder="e.g. 010b or 1080 TI" class="form-control input-lg" type="text">
		</div>
	</div>

	<!--Submit Button-->
	<div class="form-group">
		<label class="control-label" for="submit"></label>
		<div class="">
			<button id="submit" name="action" value="submit" type="submit" class="btn btn-primary">Submit</button>
		</div>
	</div>

	<!--Show All Button-->
	<div class="form-group">
		<label class="control-label" for="submit"></label>
		<div>
			<a id="list" class="btn btn-primary" href="/?ven=&dev=&action=submit">List All Devices</a>
		</div>
	</div>
</form>