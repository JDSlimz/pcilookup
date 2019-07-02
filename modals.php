<?php
?>
<!-- Contact Modal -->
<div id="contactModal" class="modal fade" role="dialog">
  <div class="modal-dialog">

    <!-- Modal content-->
    <div class="modal-content">
      <div class="modal-header">
        <h4 class="modal-title">Contact Us</h4>
        <button type="button" class="close" data-dismiss="modal">&times;</button>
      </div>
      <div class="modal-body">
      		<form class="form-horizontal" style="width:100%; text-align:center;">

				<!-- Name input-->
				<div class="form-group row">
					<label class="col-md-3 control-label" for="name">Name</label>  
					<div class="col-md-6">
						<input id="name" name="name" type="text" placeholder="" class="form-control input-md" required="">
					</div>
				</div>

				<!-- Email input-->
				<div class="form-group row">
					<label class="col-md-3 control-label" for="email">E-Mail</label>  
					<div class="col-md-6">
						<input id="email" name="email" type="email" placeholder="" class="form-control input-md" required="">
					</div>
				</div>

				<!-- Textarea -->
				<div class="form-group row">
				  <label class="col-md-3 control-label" for="message">Message</label>
				  <div class="col-md-6">                     
				    <textarea class="form-control" id="message" name="message">What do you have to say?</textarea>
				  </div>
				</div>

				<!-- CAPTCHA -->
				<script src="https://www.google.com/recaptcha/api.js?render=6Ldbm6sUAAAAAP5Bf8Griov3CARCxvz8e46B0oEJ"></script>
			    <script>
			        grecaptcha.ready(function () {
			            grecaptcha.execute('6Ldbm6sUAAAAAP5Bf8Griov3CARCxvz8e46B0oEJ', { action: 'contact' }).then(function (token) {
			                var recaptchaResponse = document.getElementById('recaptchaResponse');
			                recaptchaResponse.value = token;
			            });
			        });
			    </script>
			    <input type="hidden" name="recaptcha_response" id="recaptchaResponse">

				<!-- Button -->
				<div class="form-group row">
					<label class="col-md-3 control-label" for="submit"></label>
					<div class="col-md-6">
						<button id="message" name="action" value="message" type="submit" class="btn btn-primary">Submit</button>
					</div>
				</div>

			</form>
      </div>
    </div>

  </div>
</div>

<!-- Help Modal -->
<div id="helpModal" class="modal fade" role="dialog">
  <div class="modal-dialog modal-lg">

    <!-- Modal content-->
    <div class="modal-content">
      <div class="modal-header">
        <h4 class="modal-title">Help</h4>
        <button type="button" class="close" data-dismiss="modal">&times;</button>
      </div>
      <div class="modal-body">

      		<p>PCI Lookup is desinged to help you find the Vendor and Device descriptions you need to get drivers for you PC. If you are not sure where to start, there is some helpful information below that can get you started. If you are still lost, feel free to contact us, we would be happy to help!</p>
      		
			<ul class="nav nav-tabs">
				<li class="nav-item">
					<a class="nav-link active" data-toggle="tab" href="#windows">Windows</a>
				</li>
				<li class="nav-item">
					<a class="nav-link" data-toggle="tab" href="#linux">Linux</a>
				</li>
			</ul>

			<div class="tab-content" style="text-align: left;">
				<div class="tab-pane active container" id="windows">
					<ol>
					<li>To find the Vendor and Device IDs in Windows, you must first open your Device Manager, there are several ways to do this:</li>
						<ul>
							<li>Open the Start Menu (Windows Menu, bottom right corner of screen), simply type "Device Manager" and select it when it appears.
						</ul>
					</li>
					<li>Once you have opened the Device Manager, you need to select the device you need drivers for. It will usually have a yellow triangle with an exclamation point in it, or be called "Unknown Device". </li>

					<li>Right-click on the device in question and select "Properties", a new window will appear.</li>

					<li> In that window, click on the "Details" tab.</li>

					<li>Using the drop-down menu, select "Hardware Ids"</li>

					<li>You will see something like "PCI\VEN_8086&DEV_15B7&SUBSYS_06E01028&REV_31"
						<ul>
							<li>VEN_8086 means that the Vendors ID is 8086, you can search this in the Vendor box on the home page.</li>
							<li>DEV_1587 means that the Device ID is 1587, you can search for this using the Device box on the home page.</li>
							<li>SUBSYS_06E01028 means that the Susbsystem ID is 05E01028, you can search for this using the Device box on the home page as well.</li>
						</ul>
					</li>

					</ol>

				</div>
				<div class="tab-pane container" id="linux">
					Detailed walkthroughs can be found on Wikis and Knowledgebases below:
					<ul>
						<li><a target="_blank" href="https://wiki.debian.org/HowToIdentifyADevice/PCI">Debian</a></li>
						<li><a target="_blank" href="https://access.redhat.com/solutions/56081">Red Hat Enterprise</a></li>
					</ul>
				</div>
			</div>

      </div>
    </div>

  </div>
</div>

<!-- Bitcoin Modal -->
<div id="btcModal" class="modal fade" role="dialog">
  <div class="modal-dialog">

    <!-- Modal content-->
    <div class="modal-content">
      <div class="modal-header">
        <h4 class="modal-title">Donate Bitcoin</h4>
        <button type="button" class="close" data-dismiss="modal">&times;</button>
      </div>
      <div class="modal-body" style="text-align: center;">
        <img src="/img/btcqr.png" />
        <p class="address">1FXCiDcZKswWw1j99tKoTANDqUMjv6TUYR</p>
      </div>
      <div class="modal-footer">
        <button type="button" class="btn btn-default" data-dismiss="modal">Close</button>
      </div>
    </div>

  </div>
</div>

<!-- Litecoin Modal -->
<div id="ltcModal" class="modal fade" role="dialog">
  <div class="modal-dialog">

    <!-- Modal content-->
    <div class="modal-content">
      <div class="modal-header">
        <h4 class="modal-title">Donate Litecoin</h4>
        <button type="button" class="close" data-dismiss="modal">&times;</button>
      </div>
      <div class="modal-body" style="text-align: center;">
      	<img src="/img/ltcqr.png" />
        <p class="address">LM4sL1m1YLoMMUnVptcHRrdV6mu96XPNxS</p>
      </div>
      <div class="modal-footer">
        <button type="button" class="btn btn-default" data-dismiss="modal">Close</button>
      </div>
    </div>

  </div>
</div>

<!-- Iota Modal -->
<div id="iotaModal" class="modal fade" role="dialog">
  <div class="modal-dialog">

    <!-- Modal content-->
    <div class="modal-content">
      <div class="modal-header">
        <h4 class="modal-title">Donate Iota</h4>
        <button type="button" class="close" data-dismiss="modal">&times;</button>
      </div>
      <div class="modal-body" style="text-align: center;">
      	<img src="/img/iotaqr.png" />
        <p class="address">UXVJIHQVOMLILKXVPNOHBQGVEOPMXOFVKJSNCJAWXSNJH9PYFORHBDKSZAJMUC9CSGBO9ISP9RYZEBMIBAAVTBO9JD</p>
      </div>
      <div class="modal-footer">
        <button type="button" class="btn btn-default" data-dismiss="modal">Close</button>
      </div>
    </div>

  </div>
</div>

<!-- Nano Modal -->
<div id="nanoModal" class="modal fade" role="dialog">
  <div class="modal-dialog">

    <!-- Modal content-->
    <div class="modal-content">
      <div class="modal-header">
        <h4 class="modal-title">Donate Nano</h4>
        <button type="button" class="close" data-dismiss="modal">&times;</button>
      </div>
      <div class="modal-body" style="text-align: center;">
      	<img src="/img/nanoqr.png" />
        <p class="address">xrb_3n3tdajap6cw83nspz1uj4t4cc1wmgzd8yxpbd85bfga5xb58gwpayt3hix9</p>
      </div>
      <div class="modal-footer">
        <button type="button" class="btn btn-default" data-dismiss="modal">Close</button>
      </div>
    </div>

  </div>
</div>

<!-- Survey Modal -->
<div id="surveyModal" class="modal fade" role="dialog">
  <div class="modal-dialog modal-lg">

    <!-- Modal content-->
    <div class="modal-content">
      <div class="modal-header">
        <h4 class="modal-title">Survey!</h4>
        <button type="button" class="close" data-dismiss="modal">&times;</button>
      </div>
      <div class="modal-body">
			<p style="width:90%; text-align:center; margin: 10px auto;">Thank you for taking our short survey! Your input will help us better serve you!</p>
      		<form id="surveyForm" class="form-horizontal needs-validation" style="width:100%; text-align:center;">
			  <div id="ableToFindContainer" class="form-group row">
				<label for="ableToFind" class="control-label col-sm-6">Were you able to find your device?</label> 
				<div class="col-sm-6">
				  <label class="radio-inline">
					<input class="form-control" name="ableToFind" value="yes" type="radio">
						  Yes
				  </label>
				  <label class="radio-inline">
					<input class="form-control" name="ableToFind" value="no" type="radio">
						  No
				  </label>
				</div>
			  </div>
			  <div class="form-group row">
				<label for="deviceInfo" class="control-label col-sm-6">If not, what is the devices information?</label> 
				<div class="col-sm-6">
				  <input id="deviceInfo" name="deviceInfo" class="form-control" type="text">
				</div>
			  </div>
			  <div id="overallRatingContainer" class="form-group row">
				<label for="overallRating" class="control-label col-sm-6">Overall Rating</label> 
				<div class="col-sm-6">
				  <label class="radio-inline">
					<input class="form-control" name="overallRating" value="1" type="radio">
						  1
				  </label>
				  <label class="radio-inline">
					<input class="form-control" name="overallRating" value="2" type="radio">
						  2
				  </label>
				  <label class="radio-inline">
					<input class="form-control" name="overallRating" value="3" type="radio">
						  3
				  </label>
				  <label class="radio-inline">
					<input class="form-control" name="overallRating" value="4" type="radio">
						  4
				  </label>
				  <label class="radio-inline">
					<input class="form-control" name="overallRating" value="5" type="radio">
						  5
				  </label>
				  <label class="radio-inline">
					<input class="form-control" name="overallRating" value="6" type="radio">
						  6
				  </label>
				  <label class="radio-inline">
					<input class="form-control" name="overallRating" value="7" type="radio">
						  7
				  </label>
				  <label class="radio-inline">
					<input class="form-control" name="overallRating" value="8" type="radio">
						  8
				  </label>
				  <label class="radio-inline">
					<input class="form-control" name="overallRating" value="9" type="radio">
						  9
				  </label>
				  <label class="radio-inline">
					<input class="form-control" name="overallRating" value="10" type="radio">
						  10
				  </label>
				</div>
			  </div>
			  <div class="form-group row">
				<label for="comments" class="control-label col-sm-6">Comments, suggestions, or requests:</label> 
				<div class="col-sm-6">
				  <textarea id="comments" name="comments" cols="40" rows="5" class="form-control"></textarea>
				</div>
			  </div> 
			  <div class="form-group row">
				<div class="col-xs-offset-4 col-xs-8" style="margin: 0 auto;">
				  <input id="action" name="action" type="hidden" value="survey">
				  <button name="submit" type="button"  onClick="submitSurvey()" class="btn btn-primary">Submit</button>
				</div>
			  </div>
			</form>
      </div>
    </div>

  </div>
</div>