/*
 |--------------------------------------------------------------------------
 | Mix Asset Management
 |--------------------------------------------------------------------------
 |
 | Mix provides a clean, fluent API for defining some Webpack build steps
 | for your Laravel application. By default, we are compiling the Sass
 | file for the application as well as bundling up all the JS files.
 |
 */

/**
 * Require & Import
 * 
 */
const mix = require('laravel-mix');
const Path = require('path');
const BrowsersSupport = require("./config/browserslistrc.js");
const BrowserSyncPlugin = require('browser-sync-webpack-plugin');
//require('laravel-mix-imagemin');

/**
 * Set paths
 * 
 */
mix.setPublicPath('assets')

/**
 * Javascript
 * Compiles all JS to assets
 * 
 */
mix.js('src/js/app.js', 'assets/js/app.js')

/**
 * SCSS
 * Compiles all SCSS files to assets and uses Sass lint.
 * 
 */
mix.sass('src/scss/app.scss', 'assets/css/app.css')

/**
 * BrowserSync //! Optional
 * 
 */
// mix.browserSync({
//     proxy:"localhost:8000",
//     baseDir: "./assets/",
//     open: true,
//     notify: false
// });
// mix.webpackConfig({
//     plugins: [
//         new BrowserSyncPlugin({
//             host: 'localhost',
//             port: 8000,
//             server: { baseDir: ['assets'] },
//             files: [
//                 '**/*.css'
//             ]
//         })
//     ]
// })



/**
 * Production
 * Uses prettier plugin, minifies & babel for JS.
 * Also adds version to assets.
 * 
 */
if (mix.config.production) {

    mix.babel('assets/js/app.js', 'assets/js/app.js')

    mix.options({
        autoprefixer: {
            options: {
                overrideBrowserslist: BrowsersSupport.overrideBrowserslist
            }
        }
    })

    /**
     * Images 
     * Copies & compresses image files.
     * 
     */
    mix.imagemin(
        'img/**.*',
        {
            context: 'resources',
        },
        {
            optipng: {
                optimizationLevel: 5
            },
            jpegtran: null,
            plugins: [
                require('imagemin-mozjpeg')({
                    quality: 100,
                    progressive: true,
                }),
            ],
        }
    );
}