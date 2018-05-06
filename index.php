<?php
require_once('head.php');
require_once('header.php');

if(isset($_GET['action']) && $_GET['action'] == "submit"){
	require_once('resultsTable.php');
} else if(isset($_GET['action']) && $_GET['action'] == "message"){

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
	  require_once('form.php');
	}
} else {
	require_once('form.php');
}

require_once('modals.php');
require_once('foot.php');
?>