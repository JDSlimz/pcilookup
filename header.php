<?php
require_once('/var/www/owa/owa_php.php');
$owa = new owa_php();
$owa->setSiteId('03a075f28945bb3c70eaae6676044fc5');
$owa->setPageTitle('home');
$owa->trackPageView();
?>
<header>
	<div id="header">
		<a id="logo" href="/"><div><img src="/img/pci-text.png" /></div></a>
		<div id="menu">
			<?php
			if(isset($_GET['action']) && $_GET['action'] == "submit"){
				?> <div id="survey"><span class="oi oi-clipboard" data-toggle="modal" data-target="#surveyModal"></span></div> <?php
			}
			?>
			<div id="help"><span class="oi oi-question-mark"  data-toggle="modal" data-target="#helpModal"></span></div>
			<div id="contact"><span class="oi oi-envelope-closed" data-toggle="modal" data-target="#contactModal"></span></div>
		</div>
		<?php 
			if(isset($_GET['action']) && $_GET['action'] == "submit"){
				require_once('headForm.php');
			}
		?>
        </div>
	</header>
