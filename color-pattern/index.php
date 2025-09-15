<?php
// Размеры сетки
$rows = 10;
$cols = 10;
$squareSize = 30;

// Функция генерации случайного цвета в hex
function randomColor() {
    return sprintf('#%02X%02X%02X', rand(0,255), rand(0,255), rand(0,255));
}

?>
<!DOCTYPE html>
<html lang="ru">
<head>
<meta charset="UTF-8" />
<title>Случайный узор</title>
<style>
  body {
    margin: 20px;
  }
  .square {
    width: <?= $squareSize ?>px;
    height: <?= $squareSize ?>px;
    float: left;
  }
  .clear {
    clear: both;
  }
</style>
</head>
<body>

<h2>Случайно сгенерированный узор</h2>
<div style="width: <?= $cols * $squareSize ?>px;">
<?php
for ($r = 0; $r < $rows; $r++) {
    for ($c = 0; $c < $cols; $c++) {
        $color = randomColor();
        echo "<div class='square' style='background-color: $color;'></div>";
    }
    echo "<div class='clear'></div>";
}
?>
</div>

</body>
</html>
