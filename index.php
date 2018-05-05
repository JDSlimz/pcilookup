<?php
require_once('head.php');
require_once('header.php');

if(isset($_GET['action']) && $_GET['action'] == "submit"){
	require_once('resultsTable.php');
} else if(isset($_GET['action']) && $_GET['action'] == "messsage"){

	$to = 'admin@pcilookup.com';
	$subject = 'PCI Lookpup Contact Form';
	$msg = $_GET['message'];
	$headers  = "From: " . $_GET['name'] . " < " . $_GET['email'] . " >\n";
    $headers .= "Cc: testsite < " . $_GET['email'] . " >\n"; 
    $headers .= "X-Sender: testsite < " . $_GET['email'] . " >\n";
    $headers .= 'X-Mailer: PHP/' . phpversion();
    $headers .= "X-Priority: 1\n"; // Urgent message!
    $headers .= "Return-Path: " . $_GET['email'] . "\n"; // Return path for errors
    $headers .= "MIME-Version: 1.0\r\n";
    $headers .= "Content-Type: text/html; charset=iso-8859-1\n";

	mail($to, $subject, $msg, $headers);

} else {
	require_once('form.php');
}

require_once('modals.php');
require_once('foot.php');
?>