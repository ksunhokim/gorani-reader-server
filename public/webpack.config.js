var path = require("path");
var webpack = require("webpack");

var nodeModulesPath = path.join(__dirname, 'node_modules');
var isProduction = process.env.NODE_ENV == "production";

var config = {
  context: path.join(__dirname, 'src'),

  entry: './index.tsx',

  resolve: {
    extensions: ['.tsx', '.ts', '.js', '.scss', '.css'],
  },

  output: {
      path: path.join(__dirname, 'dist'),
      filename: 'bundle.js',
      publicPath: '/'
  },

  module: {
    rules: [
      {
        test: /\.tsx?$/,
        enforce: 'pre',
        loader: 'tslint-loader',
        options: { emitErrors: true }
      },
      {
        test: /\.tsx?$/,
        loaders: ["babel-loader?cacheDirectory", "awesome-typescript-loader?tsconfig=tsconfig.webpack.json&useCache=true"]
      },
      {
        test: /\.css$/,
        loaders: ["style-loader", "css-loader?minimize"]
      },
      {
        test: /(\.scss$)/,
        loaders: [{
            loader: 'style-loader'
        }, {
            loader: 'css-loader'
        }, {
            loader: 'sass-loader',
            options: {
                outputStyle: 'compressed',
                includePaths: ['./node_modules']
            }
        }]
      },
      {
        test: /\.(jpg|png|woff|eot|ttf|svg|gif)$/,
        loader: "file-loader?name=[name]_[hash].[ext]"
      }
    ]
  },

  plugins: [
    new webpack.DefinePlugin({
      DEBUG: true
    })
  ]
};

if (isProduction) {
  config.plugins.push(new webpack.optimize.UglifyJsPlugin({
     compress: {
        warnings: false
    }
  }));
  config.plugins.push(new webpack.DefinePlugin({
    'process.env': {NODE_ENV: '"production"'},
    DEBUG: false
  }));
}

module.exports = config;
