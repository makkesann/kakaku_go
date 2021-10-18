## このアプリケーションについて
店舗版価格.comの様なもので、食料品(現在は飲み物のみ)の実店舗の価格を調べられるサービスを制作しました。

制作の背景としては、食料品店で買い物をしている際、「この品物はどこが一番安いんだろう」と思うことが多々あり、その度に実際に足を運んで回り、大変な思いをしていたことがあります。

そういった体験から、品物ごとに安い店を瞬時に調べられるサービスがあると便利だという考えから、本サービスの制作をしました。
  
## 使用技術、機能
  
 クライアントサイド：Vue.js, scss サーバーサイド：gin(golang)　RDBMS：PostgreSQL O/Rマッパー： GORM　サーバー：Amazon EC2, S3 外部API:楽天市場商品検索API, Google Map API その他：Docker
 
  *製品追加、削除、編集機能（管理者のみ）
  
  *店舗追加、削除、編集機能（管理者のみ）
  
  *価格情報の追加
  
  *価格情報の削除、編集機能（管理者のみ）
  
  *お気に入り機能（保存は一般ユーザーログイン時のみ）
  
  *飲み物検索機能
  
  *製品一覧ページでのページネーション機能
  
  *各フォームでのバリデーション機能
  
  *レスポンシブデザイン対応
 
## 使い方について
### テストアカウント
#### 一般ユーザー
ユーザー名： test , パスワード： pass
#### 管理者
ユーザー名： admin , パスワード：s6da5g454sdf

トップページ(/drink)は飲み物一覧ページです。このページの左側にあるものはジャンルでクリックするとそのジャンルに該当する商品のみ表示されるようになります。

各製品名をクリックすると、個別ページに飛び、その製品がどの店舗でどの価格で売られているかを表示し、さらに楽天での販売価格、各店舗が近くでどこにあるかをGoogle Maps APIで表示するというものです。

ログインすると商品や店舗をお気に入り登録の状態を紐づけすることが出来、また、ログイン情報はリロードしても維持されます。

また、各商品の画像は、画像パスをデータベースに保存し、画像ファイルはS3に保存することによってアクセスすることとしました。
