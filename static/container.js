'use strict';

var _createClass = function () { function defineProperties(target, props) { for (var i = 0; i < props.length; i++) { var descriptor = props[i]; descriptor.enumerable = descriptor.enumerable || false; descriptor.configurable = true; if ("value" in descriptor) descriptor.writable = true; Object.defineProperty(target, descriptor.key, descriptor); } } return function (Constructor, protoProps, staticProps) { if (protoProps) defineProperties(Constructor.prototype, protoProps); if (staticProps) defineProperties(Constructor, staticProps); return Constructor; }; }();

function _classCallCheck(instance, Constructor) { if (!(instance instanceof Constructor)) { throw new TypeError("Cannot call a class as a function"); } }

function _possibleConstructorReturn(self, call) { if (!self) { throw new ReferenceError("this hasn't been initialised - super() hasn't been called"); } return call && (typeof call === "object" || typeof call === "function") ? call : self; }

function _inherits(subClass, superClass) { if (typeof superClass !== "function" && superClass !== null) { throw new TypeError("Super expression must either be null or a function, not " + typeof superClass); } subClass.prototype = Object.create(superClass && superClass.prototype, { constructor: { value: subClass, enumerable: false, writable: true, configurable: true } }); if (superClass) Object.setPrototypeOf ? Object.setPrototypeOf(subClass, superClass) : subClass.__proto__ = superClass; }

import { Login } from "./login.js";
import { Main } from "./main.js";

var e = React.createElement;

var Container = function (_React$Component) {
    _inherits(Container, _React$Component);

    function Container(props) {
        _classCallCheck(this, Container);

        var _this = _possibleConstructorReturn(this, (Container.__proto__ || Object.getPrototypeOf(Container)).call(this, props));

        _this.state = {
            title: "第八季第一期",
            token: "asddas",
            deadLine: "2020-2-24 19:30:00"
        };
        return _this;
    }

    _createClass(Container, [{
        key: "render",
        value: function render() {
            console.log(this.state);
            return React.createElement(
                "div",
                null,
                React.createElement(
                    "div",
                    { className: "t" },
                    React.createElement(
                        "p",
                        null,
                        "\u9752\u5E74\u5927\u5B66\u4E60"
                    )
                ),
                React.createElement(
                    "div",
                    { className: "cTitle" },
                    React.createElement(
                        "p",
                        null,
                        this.state.title
                    )
                ),
                React.createElement(
                    "div",
                    { className: "cMain" },
                    React.createElement(
                        "div",
                        { className: "cTopContain" },
                        React.createElement(
                            "p",
                            null,
                            "\u5C06\u4E8E ",
                            this.state.deadLine,
                            " \u524D\u622A\u6B62\u63D0\u4EA4"
                        )
                    )
                ),
                this.state.token === "" ? React.createElement(Login, null) : React.createElement(Main, null)
            );
        }
    }]);

    return Container;
}(React.Component);

var domContainer = document.querySelector('#container');
ReactDOM.render(e(Container), domContainer);