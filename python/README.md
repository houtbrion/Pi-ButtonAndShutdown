# python版ButtonAndShutdown

Raspberry Piにシールドを追加して，ボタンを押すと，LEDを点灯させ，シャットダウンをするためのPythonスクリプト．

## 1. サポートしているpythonのバージョン
動作確認した環境はPython3です．

## 2. ハードウェア構成(ピン配置)の変更
GPIOの26番にタクトスイッチ，19番にLEDをつけることが前提になっていますが，
ピン番号を変えたい場合は，プログラムの中で定義している以下の部分を書き換えてください．

```
## GPIO ピン番号
GPIO_PIN_BTN  = 26    # ボタン : 橙

GPIO_PIN_LED  = 19    # LED : 赤
```

## 3. インストール方法
特別なライブラリは使っていないので，インストールされているpythonの環境そのままで動くはずです．

### 3.1 プログラムの配置
/usr/local/binに「buttonAndShutdown」をコピーし，chmodしておく．

### 3.2 オプション
systemdを使う場合は，以下の手順を実行．

* 「buttonAndShutdown.service」を「/etc/systemd/system」にコピー
* 「systemctl enable buttonAndShutdown.service」
* リブート



