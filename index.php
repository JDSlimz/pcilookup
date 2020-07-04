<?php
require_once('pass.php');
require_once('head.php');
require_once('header.php');

//require_once('/var/www/owa/owa_php.php');
//$owa = new owa_php();
//$owa->setSiteId('03a075f28945bb3c70eaae6676044fc5');


if(isset($_GET['action']) && $_GET['action'] == "submit"){
	require_once('resultsTable.php');
	//$owa->setPageTitle('results');
} else if(isset($_GET['action']) && $_GET['action'] == "message" && isset($_GET['recaptcha_response'])){
	//$owa->setPageTitle('message');
	// Build POST request:
	$recaptcha_url = 'https://www.google.com/recaptcha/api/siteverify';
	$recaptcha_secret = grcKey();
	$recaptcha_response = $_GET['recaptcha_response'];

	// Make and decode POST request:
	$recaptcha = file_get_contents($recaptcha_url . '?secret=' . $recaptcha_secret . '&response=' . $recaptcha_response);
	$recaptcha = json_decode($recaptcha);

	// Take action based on the score returned:
	if ($recaptcha->score >= 0.5) {
		// Verified - send email
		$to = 'josh@pcilookup.com';
		$subject = 'PCI Lookpup Contact Form';
		$msg = $_GET['name'] . "<br>" . $_GET['email'] . "<br>" . $_GET['message'];
		$headers  = "From: " . $_GET['name'] . " < " . $_GET['email'] . " >\n";
		$headers .= "X-Sender: " . $_GET['name'] . " < " . $_GET['email'] . " >\n";
		$headers .= 'X-Mailer: PHP/' . phpversion();
		$headers .= "X-Priority: 1\n"; // Urgent message!
		$headers .= "MIME-Version: 1.0\r\n";
		$headers .= "Content-Type: text/html; charset=iso-8859-1\n";

		if(mail($to, $subject, $msg, $headers))
		{
			?>
			<span id="message" style="display:block; width:100%; text-align:center; background-color:green; color:white">Message sent! Thank you!</span>
			<script>
				setTimeout(fade_out, 5000);

				function fade_out() {
				  $("#message").fadeOut().empty();
				}
			</script>
			<?php
		  require_once('form.php');
		}else{
			?>
			<span id="message" style="display:block; width:100%; text-align:center; background-color:red; color:white">There was an error, please try again or email us directly at admin@pcilookup.com</span>
			<script>
				setTimeout(fade_out, 5000);

				function fade_out() {
				  $("#message").fadeOut().empty();
				}
			</script>
			<?php
		}
	} else {
		// Not verified - show form error
	}
	require_once('form.php');
} else {
	//$owa->setPageTitle('home');
	require_once('form.php');
}
//$owa->trackPageView();
require_once('modals.php');
require_once('foot.php');
?>