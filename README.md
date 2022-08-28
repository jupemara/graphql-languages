# graphql-languages

try out how great is graphql developer experience with each language

## 目的

- 次のプロジェクトで graphql を採用することは決まってるんだが、nestjs を触ってみた感じ、decorator でガシガシ書いていく感じが若干肌に合わない感じと、コードファーストは本当に小規模開発にあっているのかということを自分の肌感で感じたい
  - んまぁつまり TypeScript or golang どっちのほうが自分にとっていいかってのをここで示していきたい
- golang と TypeScript ( mainly nestjs ) の開発体験の違いについてまとめたい。

## 実装前の先入観 ( ネットで調べただけの知識 )

- graphql
  - controller ( graphql 的には resolver ) のレイヤーだけ graphql にかぶせたいっていうユースケースのとき、スキーマファーストとコードファーストがバッティングする恐れ
    - e.g: controller を cli でさくっとみたいなとき、スキーマファーストだと graphql のみって感じがする
      - もはやそういうときはマイクロサービス的には別の usecase (分析用途とか) になるケースなんかもしれんな。つまり特に気にしなくていいことなんかもしれん
- nestjs
  - decorator きもない?? 明示的じゃない感じがしてちょっとアレ...
  - 普通に DI したいだけやのに大げさちゃうか??
  - rails の再来感
  - `nest g s` みたいなコマンドの yoeman 感
  - typeorm 前提の DB
    - firestore とか spanner 使いたいよねってケースのときはどないするん??
  - angular に強く inspire されてるって話やけど angular の service 不要論 とかを思い出すとアーキテクチャやディレクトリ構造に柔軟性をもたせられないのでは??
  - フレームワークにのっかればめんどくさい実装はだいたい解決できる(かも?)
  - 大規模開発の各マイクロサービスチームの開発速度向上とうかコリオグラフィーを向上させるのにはとてもいいと思う
    - メルカリとか見てるとそう思う
- gqlgen
  - 見た感じかなり薄そうに見える
    - controller レイヤーだけ graphql にする的な
  - コードファーストどうやるんやろ?感

https://blog.arashike.com/nestjs-vs-gqlgen こっちにまとめたよ
