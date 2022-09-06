#!/bin/bash
#作用：将src中多个js源文件（压缩）合并成一个js文件放到dist目录中。因为这样解析起来比较快，且html中只需要引用这一个js就够了

JS_PATH=/home/zahlenw2/goapp/static/js/
JS_PATH_DIST=${JS_PATH}dist/
JS_PATH_SRC=${JS_PATH}src/

find $JS_PATH_SRC -type f -name '*.js' | sort | xargs cat > ${JS_PATH_DIST}game.js

