# Pi-ButtonAndShutdown

Raspberry Piにシールドを追加して，ボタンを押すと，LEDを点灯させ，シャットダウンをするためのPythonスクリプト．

## 1. サポートしているOSのバージョン
動作確認した環境は「Raspbian stretch」です．

## 2. ハードウェア構成
GPIOの26番にタクトスイッチ，19番にLEDをつけることが前提になっていますが，
ピン番号を変えたい場合は，プログラムの中で定義している以下の部分を書き換えてください．

```
## GPIO ピン番号
GPIO_PIN_BTN  = 26    # ボタン : 橙

GPIO_PIN_LED  = 19    # LED : 赤
```


### 2.1 ブレッドボードでの配線
![配線イメージ][breadboard]


### 2.2 回路図
![回路図][circuit]


## 3. インストール方法
特別なライブラリは使っていないので，インストールされているpythonの環境そのままで動くはず．

### 3.1 プログラムの配置
/usr/local/binに「buttonAndShutdown」をコピーし，chmodしておく．

### 3.2 オプション
systemdを使う場合は，以下の手順を実行．

* 「buttonAndShutdown.service」を「/etc/systemd/system」にコピー
* 「systemctl enable buttonAndShutdown.service」
* リブート

## 4. ライセンス
本当はライブラリ化してLGPLにするか，BSDやMITライセンスにするところだけど迷い中．とりあえずは放置．
これにしてほしいとかいう意見があればメールでもくださいな．

[breadboard]: ブレッドボード.png "配線イメージ"
[circuit]: 回路図.png "回路図"

