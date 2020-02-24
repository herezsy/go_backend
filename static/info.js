'use strict';

var _createClass = function () { function defineProperties(target, props) { for (var i = 0; i < props.length; i++) { var descriptor = props[i]; descriptor.enumerable = descriptor.enumerable || false; descriptor.configurable = true; if ("value" in descriptor) descriptor.writable = true; Object.defineProperty(target, descriptor.key, descriptor); } } return function (Constructor, protoProps, staticProps) { if (protoProps) defineProperties(Constructor.prototype, protoProps); if (staticProps) defineProperties(Constructor, staticProps); return Constructor; }; }();

function _classCallCheck(instance, Constructor) { if (!(instance instanceof Constructor)) { throw new TypeError("Cannot call a class as a function"); } }

function _possibleConstructorReturn(self, call) { if (!self) { throw new ReferenceError("this hasn't been initialised - super() hasn't been called"); } return call && (typeof call === "object" || typeof call === "function") ? call : self; }

function _inherits(subClass, superClass) { if (typeof superClass !== "function" && superClass !== null) { throw new TypeError("Super expression must either be null or a function, not " + typeof superClass); } subClass.prototype = Object.create(superClass && superClass.prototype, { constructor: { value: subClass, enumerable: false, writable: true, configurable: true } }); if (superClass) Object.setPrototypeOf ? Object.setPrototypeOf(subClass, superClass) : subClass.__proto__ = superClass; }

var e = React.createElement;

var Info = function (_React$Component) {
    _inherits(Info, _React$Component);

    function Info(props) {
        _classCallCheck(this, Info);

        var _this = _possibleConstructorReturn(this, (Info.__proto__ || Object.getPrototypeOf(Info)).call(this, props));

        _this.state = {
            name: "*舜宇",
            stuid: 201726012601,
            usePassword: false
        };
        return _this;
    }

    _createClass(Info, [{
        key: "render",
        value: function render() {
            return React.createElement(
                "div",
                { className: "info" },
                React.createElement(
                    "div",
                    { className: "cInfo" },
                    React.createElement(
                        "p",
                        null,
                        this.state.name
                    ),
                    React.createElement(
                        "p",
                        null,
                        this.state.stuid
                    )
                ),
                React.createElement(
                    "div",
                    {
                        className: "cInfo"
                    },
                    React.createElement(
                        "p",
                        null,
                        this.state.usePassword ? "修改密码" : "设置密码"
                    ),
                    React.createElement(
                        "p",
                        null,
                        "\u9000\u51FA"
                    )
                )
            );
        }
    }]);

    return Info;
}(React.Component);

export { Info };