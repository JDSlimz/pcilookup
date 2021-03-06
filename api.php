<?php
require_once('pass.php');
//Declare and define variables
$action = "";
$vendor = "";
$device = "";
$pass = "";

//Variables for survey
$ableToFind = "";
$deviceInfo = "";
$overallRating = "";
$comments = "";

class Device {
	public $id;
	public $desc;
	public $venID;
	public $venDesc;

	public function __construct($id, $desc, $venID, $venDesc) {
		$this->id = $id;
		$this->desc = $desc;
		$this->venID = $venID;
		$this->venDesc = $venDesc;
	}
}

if(isset($_GET['action'])){
	$action = $_GET['action'];
}

if(isset($_GET['pass'])){
	$pass = $_GET['pass'];
}

//Connect to database
$servername = "localhost:3306";
$username = "pci";
$password = mysqlPass();

// Create connection
$conn = new mysqli($servername, $username, $password, "pci");

if(isset($_GET['vendor']) && $_GET['vendor'] != ""){
	$vendor = mysqli_real_escape_string($conn, $_GET['vendor']);
}

if(isset($_GET['device']) && $_GET['device'] != ""){
	$device = mysqli_real_escape_string($conn, $_GET['device']);
}

// Check connection
if ($conn->connect_error) {
    die("Connection failed: " . $conn->connect_error);
}

function searchList($con, $vendor, $device){
	$devices = [];
	$sql = "SELECT id,descrip,venid,vendesc FROM devices WHERE (venid LIKE '%$vendor%' OR vendesc LIKE '%$vendor%') AND (id LIKE '%$device%' OR descrip LIKE '%$device%')";
	if($vendor == "" && $device == ""){
		$sql = "SELECT id,descrip,venid,vendesc FROM devices;";
	}
	if ($result=mysqli_query($con,$sql)){
		// Fetch one and one row
		while ($row=mysqli_fetch_row($result)){
			$device = new Device($row[0], $row[1], $row[2], $row[3]);
			array_push($devices, $device);
		}
		// Free result set
		mysqli_free_result($result);
	}
	mysqli_close($con);
	echo json_encode($devices);
}

function readFileAndUpdate($con, $url){
	set_time_limit(300) ;

	$cSession = curl_init();
	//step2
	curl_setopt($cSession,CURLOPT_URL,$url);
	curl_setopt($cSession,CURLOPT_RETURNTRANSFER,true);
	curl_setopt($cSession,CURLOPT_HEADER, false);
	//step3
	$result=curl_exec($cSession);
	//step4
	curl_close($cSession);
	//step5
	$result = explode("\n", $result);

	$thisVendor = "";
	$devices = [];

	foreach ($result as $line) {
		if (substr($line, 0, 1) != '#' && $line != "") {
			if(strspn($line, "\t") === 0){
				//Line is a vendor, save for future use!
				$thisVendor = explode(' ', $line, 2);
			} else if(strspn($line, "\t") > 0){
				//Line is a device, get vendor and create object!
				if(strspn($line, "\t") === 1){
					//device ID is 4 characters
					$thisDevice = explode(' ', $line, 2);
					$deviceID = $thisDevice[0];
					$deviceDesc = $thisDevice[1];
					$vendorID = $thisVendor[0];
					$vendorDesc = $thisVendor[1];

					$device = new Device(trim(preg_replace('/\s+/S', '',$deviceID)), $deviceDesc, $vendorID, $vendorDesc);
					array_push($devices, $device);
				} else if(strspn($line, "\t") === 2){
					//device ID is 8 characters
					$thisDevice = explode(' ', $line, 3);
					$deviceID = $thisDevice[0] . " " . $thisDevice[1];
					$deviceDesc = $thisDevice[2];
					$vendorID = $thisVendor[0];
					$vendorDesc = $thisVendor[1];

					$device = new Device(trim(preg_replace('/\s+/S', '',$deviceID)), $deviceDesc, $vendorID, $vendorDesc);
					array_push($devices, $device);
				}
			}
		}
	}

	echo count($devices) . "<br>";

	$query = "INSERT INTO devices (id, descrip, venid, vendesc) VALUES (?, ?, ?, ?)";//ON DUPLICATE KEY UPDATE descrip=?
	$stmt = $con->prepare($query);

	$con->query("START TRANSACTION");
	$count = 0;
	foreach ($devices as $device) {
		$stmt->bind_param("ssss", $device->id, $device->desc, $device->venID, $device->venDesc);//, $device->desc
		$stmt->execute();
		$count++;
	}

	echo $count . "<br>";
	$stmt->close();
	$con->query("COMMIT");


}

function submitSurvey($con){
	if(isset($_GET['ableToFind'])){
		$ableToFind = mysqli_real_escape_string($con, $_GET['ableToFind']);
	}
	if(isset($_GET['deviceInfo'])){
		$deviceInfo = mysqli_real_escape_string($con, $_GET['deviceInfo']);
	}
	if(isset($_GET['overallRating'])){
		$overallRating = intval($_GET['overallRating']);
	}
	if(isset($_GET['comments'])){
		$comments = mysqli_real_escape_string($con, $_GET['comments']);
	}

	$query = "INSERT INTO survey (ableToFind, deviceInfo, overallRating, comments) VALUES (?, ?, ?, ?)";
	$stmt = $con->prepare($query);
	$con->query("START TRANSACTION");
	$stmt ->bind_param("ssss", $ableToFind, $deviceInfo, $overallRating, $comments);
	$stmt->execute();
	$stmt->close();
	$con->query("COMMIT");

	$headers  = "From: " . $_GET['name'] . " < " . $_GET['email'] . " >\n";
    $headers .= "X-Sender: " . $_GET['name'] . " < " . $_GET['email'] . " >\n";
    $headers .= 'X-Mailer: PHP/' . phpversion();
    $headers .= "X-Priority: 1\n";
    $headers .= "MIME-Version: 1.0\r\n";
    $headers .= "Content-Type: text/html; charset=iso-8859-1\n";

	$to = 'josh@pcilookup.com';
	$subject = 'PCI Lookpup Survey';
	$msg = $ableToFind . " " . $deviceInfo . " " . $overallRating . " " . $comments;

	mail($to, $subject, $msg, $headers);
}

if($action == "search"){
	searchList($conn, $vendor, $device);
} else if($action == "update" && $pass == apiPass()){
	readFileAndUpdate($conn, "https://pci-ids.ucw.cz/v2.2/pci.ids");
	readFileAndUpdate($conn, "http://www.linux-usb.org/usb.ids");
} else if($action == "survey"){
	submitSurvey($conn);
}
?>
