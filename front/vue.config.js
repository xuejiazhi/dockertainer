// vue.config.js
module.exports = {
    configureWebpack: {
      plugins: [
        // new MyAwesomeWebpackPlugin()
      ],
      module:{
        rules: [{
          test: /\.scss$/,
          use: [
              {
                  loader: 'sass-loader',
              }
          ]
      }]
      }
    },
  }