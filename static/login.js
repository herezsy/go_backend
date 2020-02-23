'use strict';

var _createClass = function () { function defineProperties(target, props) { for (var i = 0; i < props.length; i++) { var descriptor = props[i]; descriptor.enumerable = descriptor.enumerable || false; descriptor.configurable = true; if ("value" in descriptor) descriptor.writable = true; Object.defineProperty(target, descriptor.key, descriptor); } } return function (Constructor, protoProps, staticProps) { if (protoProps) defineProperties(Constructor.prototype, protoProps); if (staticProps) defineProperties(Constructor, staticProps); return Constructor; }; }();

function _classCallCheck(instance, Constructor) { if (!(instance instanceof Constructor)) { throw new TypeError("Cannot call a class as a function"); } }

function _possibleConstructorReturn(self, call) { if (!self) { throw new ReferenceError("this hasn't been initialised - super() hasn't been called"); } return call && (typeof call === "object" || typeof call === "function") ? call : self; }

function _inherits(subClass, superClass) { if (typeof superClass !== "function" && superClass !== null) { throw new TypeError("Super expression must either be null or a function, not " + typeof superClass); } subClass.prototype = Object.create(superClass && superClass.prototype, { constructor: { value: subClass, enumerable: false, writable: true, configurable: true } }); if (superClass) Object.setPrototypeOf ? Object.setPrototypeOf(subClass, superClass) : subClass.__proto__ = superClass; }

var e = React.createElement;

var Login = function (_React$Component) {
    _inherits(Login, _React$Component);

    function Login(props) {
        _classCallCheck(this, Login);

        var _this = _possibleConstructorReturn(this, (Login.__proto__ || Object.getPrototypeOf(Login)).call(this, props));

        _this.state = {
            stuid: "",
            password: ""
        };
        return _this;
    }

    _createClass(Login, [{
        key: "render",
        value: function render() {
            var _this2 = this;

            return React.createElement(
                "div",
                { className: "cMain" },
                React.createElement(
                    "div",
                    { className: "cMainContain" },
                    React.createElement(
                        "div",
                        { className: "cRow" },
                        React.createElement(
                            "div",
                            { className: "cLeft" },
                            React.createElement(
                                "p",
                                null,
                                "\u5B66\u53F7"
                            )
                        ),
                        React.createElement(
                            "div",
                            { className: "cRight" },
                            React.createElement("input", {
                                className: "cInput",
                                type: "number",
                                onChange: function onChange(e) {
                                    _this2.setStatue({
                                        stuid: e.target.value
                                    });
                                }
                            })
                        )
                    ),
                    React.createElement(
                        "div",
                        { className: "cRow" },
                        React.createElement(
                            "div",
                            { className: "cLeft" },
                            React.createElement(
                                "p",
                                null,
                                "\u59D3\u540D"
                            )
                        ),
                        React.createElement(
                            "div",
                            { className: "cRight" },
                            React.createElement(
                                "p",
                                null,
                                "*\u821C\u5B87"
                            )
                        )
                    ),
                    React.createElement(
                        "div",
                        { className: "cRow" },
                        React.createElement(
                            "div",
                            { className: "cLeft" },
                            React.createElement(
                                "p",
                                null,
                                "\u5BC6\u7801"
                            )
                        ),
                        React.createElement(
                            "div",
                            { className: "cRight" },
                            React.createElement("input", {
                                className: "cInput",
                                type: "text",
                                onChange: function onChange(e) {
                                    _this2.setStatue({
                                        password: e.target.value
                                    });
                                }
                            })
                        )
                    ),
                    React.createElement(
                        "div",
                        { className: "cConfirm" },
                        React.createElement(
                            "p",
                            null,
                            "\u8FDB\u5165"
                        )
                    )
                )
            );
        }
    }]);

    return Login;
}(React.Component);

export { Login };