<?php

?>
<header>
	<div id="header">
		<a id="logo" href="/"><div><img src="/img/pci-text.png" /></div></a>
		<div id="menu">
			<div id="help"><img src="/img/help.png" data-toggle="modal" data-target="#helpModal"/></div>
			<div id="contact"><img src="/img/mail-blue.png" data-toggle="modal" data-target="#contactModal"/></div>
		</div>
		<?php 
			if(isset($_GET['action']) && $_GET['action'] == "submit"){
				require_once('headForm.php');
			}
		?>
	</div>
</header>