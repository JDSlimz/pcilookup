<?php
require_once('head.php');
require_once('header.php');

if(isset($_GET['action']) && $_GET['action'] == "submit"){
	require_once('resultsTable.php');
} else if(isset($_GET['action']) && $_GET['action'] == "update"){

} else {
	require_once('form.php');
}

require_once('modals.php');
require_once('foot.php');
?>