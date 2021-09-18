# go版ButtonAndShutdown

Python版と比較すると，configファイルで回路のピン番号の変更や，logの出力方法(syslogを使うか否かなど)を選択できるように機能を強化したものになってます．

## 1. 動作環境

### 1.1 動作を確認した各種ソフトのバージョン
- go言語 : go version go1.17 linux/arm
- periph
  - periph.io/x/conn/v3 : v3.6.8
  - periph.io/x/d2xx    : v0.0.1
  - periph.io/x/host/v3 : v3.7.0

### 1.2 ビルドに必要なその他のツール
- make


## 2. デバイスを接続するピン番号の変更

GPIOのピン番号は，設定ファイル(本リポジトリのbuttonAndShutdown.cfg)で変更可能です．

以下は例であるが，スイッチ(ボタン)のGPIOの番号は「buttonpin」の行で設定しており，GPIO番号20番(ピン番号38番:[参考資料][gpio_map])にしている．同じくLEDは「ledpin」の行でGPIO番号16番(ピン番号36番)にしています．
```
$ cat buttonAndShutdown.cfg
{
        "usesyslog":true,
        "usestdout":false,
        "logfilename":"",
        "buttonpin":"GPIO20",
        "ledpin":"GPIO16"
}
$
```

## 3. インストール
### 3.1 準備
1章のリストを見て，go言語とmakeだけはインストールしておいてください．goやmakeをインストールする際に，その他の開発ツールを入れる必要があると思いますが，それについては，go言語の[オフィシャルサイト][golang]の指示に従ってください．

### 3.2 調整
#### 3.2.1 設定ファイルの内容を回路に合わせる

2.3節の「デバイスを接続するピン番号の変更」でも説明したように，回路に合わせて使っているデバイスを接続したGPIOの番号に合わせて同封の「buttonAndShutdown.cfg」の「buttonpin」と「ledpin」の行を変更してください．


#### 3.2.2 ログ出力先の選択
再度の掲示になりますが，設定ファイルの中身は以下のようになっています．ここで，「usesyslog」はsyslogを使うか否か，「usestdout」は標準出力にログを垂れ流す場合(回路の動作確認などに利用)，「logfilename」にログファイルのフルパスを指定すると，指定したファイルにログを追加書き込みしていきます．
```
$ cat buttonAndShutdown.cfg
{
        "usesyslog":true,
        "usestdout":false,
        "logfilename":"",
        "buttonpin":"GPIO20",
        "ledpin":"GPIO16"
}
$
```
複数のログ出力先を有効にすることも可能ですし，全部ON(もしくはOFF)にするものOKです．

#### 3.2.3 インストール先ディレクトリの修正
現状インストール先のディレクトリは/usr/localになっています．
もし，修正したい場合は以下の項目を修正してください．

- Makefile : BASE_DIRの設定行(例:BASE_DIR=/usr/local)
- プログラム本体(buttonAndShutdown.go) : 設定ファイルのフルパス(例 : const defaultConfigFileName string = "/usr/local/etc/buttonAndShutdown.cfg")
- systemdの設定(buttonAndShutdown.service) : バイナリパス名(例 : ExecStart=/usr/local/bin/buttonAndShutdown)

### 3.3 インストール
以下の2ステップです．
- ```make all```
- ```sudo make install```

rebootで動作しますが，もし，即座に動かしたい場合は，以下のコマンドを実行してください．
```
# systemctl restart buttonAndShutdown.service
```
うまく動作していれば，systemctlコマンドで「Active: active (running) since 時刻」となるはずです．
```
# systemctl status buttonAndShutdown.service
● buttonAndShutdown.service - buttonAndShutdown
     Loaded: loaded (/etc/systemd/system/buttonAndShutdown.service; enabled; ve>
     Active: active (running) since Sun 2021-09-05 17:44:10 JST; 7h ago
   Main PID: 396 (buttonAndShutdo)
      Tasks: 9 (limit: 2059)
        CPU: 13.664s
     CGroup: /system.slice/buttonAndShutdown.service
             mq396 /usr/local/bin/buttonAndShutdown

Sep 05 17:44:10 raspberrypi systemd[1]: Started buttonAndShutdown.
#
```

[gpio_map]: https://www.ishikawa-lab.com/RasPi_index.html "GPIOの参考資料"
[golang]: https://go.dev/ "go公式サイト"
[fritzing]: https://fritzing.org/ "fritzing"
