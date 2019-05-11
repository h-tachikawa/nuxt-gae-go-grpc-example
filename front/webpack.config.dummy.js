//WebStormから~や@の補完を効かせるためのダミーファイルなので、このファイルにwebpackの設定を追加しても意味がないことに注意
//Language and Framework->JavaScript→webpackの設定ファイルとしてこのファイルを指定すると補完が効くようになる

const path = require('path')

module.exports = {
  resolve: {
    extensions: ['.js', '.json', '.vue', '.ts'],
    root: path.resolve(__dirname),
    alias: {
      '@': path.resolve(__dirname),
      '~': path.resolve(__dirname),
    },
  },
}
