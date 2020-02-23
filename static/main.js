'use strict';

var _createClass = function () { function defineProperties(target, props) { for (var i = 0; i < props.length; i++) { var descriptor = props[i]; descriptor.enumerable = descriptor.enumerable || false; descriptor.configurable = true; if ("value" in descriptor) descriptor.writable = true; Object.defineProperty(target, descriptor.key, descriptor); } } return function (Constructor, protoProps, staticProps) { if (protoProps) defineProperties(Constructor.prototype, protoProps); if (staticProps) defineProperties(Constructor, staticProps); return Constructor; }; }();

function _classCallCheck(instance, Constructor) { if (!(instance instanceof Constructor)) { throw new TypeError("Cannot call a class as a function"); } }

function _possibleConstructorReturn(self, call) { if (!self) { throw new ReferenceError("this hasn't been initialised - super() hasn't been called"); } return call && (typeof call === "object" || typeof call === "function") ? call : self; }

function _inherits(subClass, superClass) { if (typeof superClass !== "function" && superClass !== null) { throw new TypeError("Super expression must either be null or a function, not " + typeof superClass); } subClass.prototype = Object.create(superClass && superClass.prototype, { constructor: { value: subClass, enumerable: false, writable: true, configurable: true } }); if (superClass) Object.setPrototypeOf ? Object.setPrototypeOf(subClass, superClass) : subClass.__proto__ = superClass; }

import { Info } from "./info.js";

var e = React.createElement;

var Main = function (_React$Component) {
    _inherits(Main, _React$Component);

    function Main(props) {
        _classCallCheck(this, Main);

        var _this = _possibleConstructorReturn(this, (Main.__proto__ || Object.getPrototypeOf(Main)).call(this, props));

        _this.state = {
            imgSrc: "./img/kp.jpg"
        };
        return _this;
    }

    _createClass(Main, [{
        key: "render",
        value: function render() {
            var imgSrc = this.state.imgSrc;
            return React.createElement(
                "div",
                null,
                React.createElement(Info, null),
                React.createElement(
                    "div",
                    { className: "cMain" },
                    React.createElement(
                        "div",
                        { className: "cMainContain" },
                        React.createElement(
                            "p",
                            { className: "cTip" },
                            imgSrc === "" ? "还未提交任何内容" : "已提交，重新上传将覆盖原内容"
                        ),
                        imgSrc === "" ? null : React.createElement("img", {
                            className: "cImg",
                            src: imgSrc,
                            alt: "\u56FE\u7247\u52A0\u8F7D\u672A\u6210\u529F"
                        }),
                        React.createElement(
                            "div",
                            { className: "cConfirm" },
                            React.createElement(
                                "p",
                                null,
                                imgSrc === "" ? "上传" : "重新上传"
                            )
                        ),
                        React.createElement(
                            "div",
                            { className: "cTip" },
                            React.createElement(
                                "p",
                                null,
                                "\u4E0A\u4F20\u540E\uFF0C\u56FE\u7247\u5C06\u88AB\u81EA\u52A8\u91CD\u547D\u540D\u81F3\u7B26\u5408\u8981\u6C42\u3002"
                            )
                        )
                    )
                )
            );
        }
    }]);

    return Main;
}(React.Component);

export { Main };