# Pi-ButtonAndShutdown

Raspberry Piにシールドを追加して，ボタンを押すと，LEDを点灯させ，シャットダウンをするためのプログラムです．
本リポジトリにはpython版とgo版がありますが，go版はconfigファイルでファンを制御するピン番号を指定できるが，pythonは
ソースを修正する必要があるなどの違いがあるため，pythonディレクトリ，goディレクトリ配下のドキュメントを見て，
好みの方を利用してください．

## 1. サポートしているOSのバージョン
動作確認した環境はRaspbianのstretchとbullseyeです．


## 2. ハードウェア構成
以下の配線図などでは，GPIOの26番にタクトスイッチ，19番にLEDをつけることが前提になっていますが，
ピン番号を変えることも可能です．プログラムにPython版を使う場合はソースを修正，go版の場合は設定ファイルの編集で
対応できます．なお，具体的な方法はgoとpythonの各ディレクトリ内のREADME.mdを参照してください．

プログラムの中で定義している以下の部分を書き換えてください．

### 2.1 ブレッドボードでの配線
![配線イメージ][breadboard]


### 2.2 回路図
![回路図][circuit]


### 2.3 CADファイル
ブレッドボードの配線イメージと回路図の元になっている
[CADソフト(Fritzing)][fritzing]の配線図もcadディレクトリに同封してあります．

## 3. インストール方法
[goディレクトリ][go]と[pythonディレクトリ][python]以下のREADME.mdを見て，どちらを使うか
判断し，使う方のディレクトリのREADME.mdの手順に従ってインストールしてください．


[breadboard]: ブレッドボード.png "配線イメージ"
[circuit]: 回路図.png "回路図"
[python]: python/README.md
[go]: go/README.md
[fritzing]: https://fritzing.org/ "fritzing"

