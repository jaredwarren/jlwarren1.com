(function($) {

    /**
     * Web Socket
     */
    var _ws;

    function _openWS(url) {
        if (_ws) {
            console.warn("Connection already Open");
            return false;
        }
        _ws = new WebSocket(url);
        _ws.onopen = function(evt) {
            _settings.onConnect(evt);
        }
        _ws.onclose = function(evt) {
            _ws = null;
            _settings.onClose(evt);
        }
        _ws.onmessage = function(evt) {
            _parseMessage(evt.data);
        }
        _ws.onerror = function(evt) {
            _settings.onError(evt);
        }
        return false;
    }

    function _closeWS() {
        if (!_ws) {
            return false;
        }
        _ws.close();
        _ws = null;
        return false;
    }

    function _sendWS(value) {
        if (!_ws) {
            return false;
        }
        _ws.send(value);
        return false;
    }


    /**
     * Scale
     */
    var _settings;

    // Parse message from websocket
    function _parseMessage(message) {
        var data = JSON.parse(message);
        if (data.steady) {
            _settings.onSteady(data);
        } else {
            _settings.onUpdate(data);
        }
    }


    var methods = {
        init: function(options) {
            // Setup default options.
            _settings = $.extend({
                url: '',
                autoConnect: true,
                // Callbacks
                onUpdate: function() {},
                onSteady: function() {},
                onConnect: function() {},
                onClose: function() {},
                onError: function() {}
            }, options);

            if (_settings.autoConnect) {
                if (!_settings.url) {
                    console.error("url option is required when autoConnect is true.");
                } else {
                    _openWS(_settings.url);
                }
            }
            return this;
        },
        connect: function(url) {
            if (typeof url == 'undefined') {
                url = _settings.url;
            }
            _openWS(url);
        },
        close: function() {
            _closeWS();
        },
        send: function(value) {
            _sendWS(value);
        }
    };


    $.fn.scale = function(methodOrOptions) {
        // Route options and/or methods
        if (methods[methodOrOptions]) {
            return methods[methodOrOptions].apply(this, Array.prototype.slice.call(arguments, 1));
        } else if (typeof methodOrOptions === 'object' || !methodOrOptions) {
            // Default to "init"
            return methods.init.apply(this, arguments);
        } else {
            $.error('Method ' + methodOrOptions + ' does not exist on jQuery.scale');
        }
    };

}(jQuery));