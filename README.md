# Pi-ButtonAndShutdown

Raspberry Piにシールドを追加して，ボタンを押すと，LEDを点灯させ，シャットダウンをするためのPythonスクリプト．


## インストール方法

### プログラムの配置
/usr/local/binに「buttonAndShutdown」をコピーし，chmodしておく．

### オプション
systemdを使う場合は，以下の手順を実行．

* 「buttonAndShutdown.service」を「/etc/systemd/system」にコピー
* 「systemctl enable buttonAndShutdown.service」
* リブート
