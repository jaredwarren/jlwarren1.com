(function($) {

    /**
     * Printer
     */

    function _getDefault(type) {
        var printer = "";
        $.ajax({
            method: 'GET',
            url: _settings.url + "/get_default/" + type,
            async: false, // must be async
            success: function(responseText, status, request) {
                printer = responseText.printer;
            },
            error: function(jqXHR, status, error) {
                console.error(jqXHR, status, error);
                _settings.onError.call(this, jqXHR, status, error);
            }
        });
        return printer;
    }

    function _connect(printerId, successCallback, errorCallback) {
        $.ajax({
            method: 'GET',
            async: _settings.async,
            url: _settings.url + "/" + printerId + '/connect',
            success: function(responseText, status, request) {
                console.info("Web Printer connected - ", responseText, status, request);
                if (typeof successCallback == "function") {
                    successCallback.call(this, responseText, status, request);
                } else {
                    _settings.onConnect.call(this, responseText, status, request);
                }
            },
            error: function(jqXHR, status, error) {
                console.error(jqXHR, status, error);
                if (typeof errorCallback == "function") {
                    errorCallback.call(this, jqXHR, status, error);
                } else {
                    _settings.onError.call(this, jqXHR, status, error);
                }
            }
        });
    }

    function _print(printerId, raw_code, job, data_type, successCallback, errorCallback) {
        if (!job) {
            job = "Web Print";
        }
        if (!data_type) {
            data_type = "RAW";
        }
        var data = {
            data: raw_code,
            job: job,
            data_type: data_type
        }
        $.ajax({
            method: 'POST',
            url: _settings.url + "/" + printerId + '/print',
            contentType: 'application/json; charset=utf-8',
            data: JSON.stringify(data),
            async: _settings.async,
            success: function(responseText, status, request) {
                console.info("Web Printer Printed - ", responseText, status, request);
                if (typeof successCallback == "function") {
                    successCallback.call(this, responseText, status, request);
                } else {
                    _settings.onPrint.call(this, responseText, status, request);
                }
            },
            error: function(jqXHR, status, error) {
                console.error(jqXHR, status, error);
                if (typeof errorCallback == "function") {
                    errorCallback.call(this, jqXHR, status, error);
                } else {
                    _settings.onError.call(this, jqXHR, status, error);
                }
            }
        });
    }

    function _printUrl(printerId, file_url, job, data_type, successCallback, errorCallback) {
        if (!job) {
            job = "Web Print";
        }
        if (!data_type) {
            data_type = "RAW";
        }
        var data = {
            file_url: file_url,
            job: job,
            data_type: data_type
        }
        $.ajax({
            method: 'POST',
            url: _settings.url + "/" + printerId + '/print_url',
            contentType: 'application/json; charset=utf-8',
            data: JSON.stringify(data),
            async: _settings.async,
            success: function(responseText, status, request) {
                console.info("Web Printer Printed - ", responseText, status, request);
                if (typeof successCallback == "function") {
                    successCallback.call(this, responseText, status, request);
                } else {
                    _settings.onPrint.call(this, responseText, status, request);
                }
            },
            error: function(jqXHR, status, error) {
                console.error(jqXHR, status, error);
                if (typeof errorCallback == "function") {
                    errorCallback.call(this, jqXHR, status, error);
                } else {
                    _settings.onError.call(this, jqXHR, status, error);
                }
            }
        });
    }

    var _settings;

    var methods = {
        init: function(options) {
            // Setup default options.
            _settings = $.extend({
                autoConnect: true,
                async: true,
                type: '', // default type of printer 2x1 or 4x6
                printerId: '', // id or name of printer
                url: '', // base url
                // Callbacks
                onPrint: function() {},
                onUpdate: function() {},
                onSteady: function() {},
                onConnect: function() {},
                onClose: function() {},
                onError: function() {}
            }, options);

            // remove trailing slash
            _settings.url = _settings.url.replace(/\/$/, "");
            if (!_settings.url) {
                console.error("Cannot connect; Missing url");
                return this;
            }
            if (_settings.autoConnect) {
                if (!_settings.printerId) {
                    if (!_settings.type) {
                        console.error("Cannot autoconnect without type or printerId");
                        return this;
                    }
                    _settings.printerId = _getDefault(_settings.type);
                }
                _connect(_settings.printerId);
            }
            return this;
        },
        connect: function(printerId, successCallback, errorCallback) {
            return _connect(printerId, successCallback, errorCallback);
        },
        close: function() {

        },
        list: function() {
            return _list();
        },
        getDefault: function(type) {
            return _getDefault(type);
        },
        print: function(printerId, raw_code, job, data_type, successCallback, errorCallback) {
            if (!printerId) {
                printerId = 'default';
            }
            if (printerId.toLowerCase() == 'default') {
                printerId = _settings.printerId;
            }
            return _print(printerId, raw_code, job, data_type, successCallback, errorCallback);
        },
        printUrl: function(printerId, file_url, job, data_type, successCallback, errorCallback) {
            if (!printerId) {
                printerId = 'default';
            }
            if (printerId.toLowerCase() == 'default') {
                printerId = _settings.printerId;
            }
            return _printUrl(printerId, file_url, job, data_type, successCallback, errorCallback);
        },
        settings: function(key, value) {
            _settings[key] = value;
        }
    };


    $.fn.printer = function(methodOrOptions) {
        // Route options and/or methods
        if (methods[methodOrOptions]) {
            return methods[methodOrOptions].apply(this, Array.prototype.slice.call(arguments, 1));
        } else if (typeof methodOrOptions === 'object' || !methodOrOptions) {
            // Default to "init"
            return methods.init.apply(this, arguments);
        } else {
            $.error('Method ' + methodOrOptions + ' does not exist on jQuery.printer');
        }
    };

}(jQuery));

/*
 * getQzRawCodeForLabel
 *
 *   Returns the raw code to send to a zebra printer for the given labelType
 *
 * Params:
 *   Same as printQzLabel
 *
 */
function getQzRawCodeForLabel(labelType) {
    // Sanitize all arguments after labelType so they can be included directly in the qz raw code string
    for (var i = 1; i < arguments.length; i++) {
        arguments[i] = cleanZPLValue(arguments[i]);
    }

    var qtyToPrint = arguments[1];
    // These code was generated by creating a layout in ZebraDesigner 2, then printing to a file, then manually escaping all newlines, backslashes, and the first character in the file
    switch (labelType.toLowerCase()) {
        case 'order_label':
            // Generate raw code to print a barcode with the order # and customer name below it
            var barcode = arguments[2];
            var subtext = arguments[3];
            return "\u0010CT~~CD,~CC^~CT~\n^XA~TA000~JSN^LT0^MNW^MTD^PON^PMN^LH0,0^JMA^PR4,4~SD15^JUS^LRN^CI0^XZ\n^XA\n^MMT\n^PW448\n^LL0153\n^LS0\n^BY2,3,61^FT41,76^BCN,,N,N\n^FD>:" + barcode + "^FS\n^FT41,104^A0N,23,24^FH\\^FD" + barcode + "^FS\n^FT41,135^A0N,23,24^FH\\^FD" + subtext + "^FS\n^PQ" + qtyToPrint + ",0,1,Y^XZ";
        case 'item_upc_label':
            // Generate raw code to print a barcode with text1 and text2 above it
            var barcode = arguments[2];
            var text1 = arguments[3];
            var text2 = arguments[4];

            // Determine the barcode type
            var barcode_type = getBarcodeType(barcode);
            if (barcode_type == 'EAN-8') {
                // Use EAN-8 layout for barcode
                return "\u0010CT~~CD,~CC^~CT~\n^XA~TA000~JSN^LT0^MNW^MTD^PON^PMN^LH0,0^JMA^PR4,4~SD15^JUS^LRN^CI0^XZ\n^XA\n^MMT\n^PW448\n^LL0153\n^LS0\n^BY2,2,53^FT20,133^B8N,,Y,N\n^FD" + barcode + "^FS\n^FT20,37^A0N,23,24^FH\\^FD" + text1 + "^FS\n^FT20,69^A0N,23,24^FH\\^FD" + text2 + "^FS\n^PQ" + qtyToPrint + ",0,1,Y^XZ";
            } else if (barcode_type == 'UPC-A') {
                // Use UPC-A layout for barcode
                return "\u0010CT~~CD,~CC^~CT~\n^XA~TA000~JSN^LT0^MNW^MTD^PON^PMN^LH0,0^JMA^PR4,4~SD15^JUS^LRN^CI0^XZ\n^XA\n^MMT\n^PW448\n^LL0153\n^LS0\n^BY2,2,53^FT38,133^BUN,,Y,N\n^FD" + barcode + "^FS\n^FT20,37^A0N,23,24^FH\\^FD" + text1 + "^FS\n^FT20,69^A0N,23,24^FH\\^FD" + text2 + "^FS\n^PQ" + qtyToPrint + ",0,1,Y^XZ";
            } else if (barcode_type == 'EAN-13') {
                // Use EAN-13 layout for barcode
                return "\u0010CT~~CD,~CC^~CT~\n^XA~TA000~JSN^LT0^MNW^MTD^PON^PMN^LH0,0^JMA^PR4,4~SD15^JUS^LRN^CI0^XZ\n^XA\n^MMT\n^PW448\n^LL0153\n^LS0\n^BY2,2,53^FT38,133^BEN,,Y,N\n^FD" + barcode + "^FS\n^FT20,37^A0N,23,24^FH\\^FD" + text1 + "^FS\n^FT20,69^A0N,23,24^FH\\^FD" + text2 + "^FS\n^PQ" + qtyToPrint + ",0,1,Y^XZ";
            } else {
                console.error("Unsupported Barcode Type", barcode_type);
            }
            // Barcode type is unknown or the checkdigit is not valid
            // Use CODE128 layout for barcode (can represent any length of letters and numbers)
            return "\u0010CT~~CD,~CC^~CT~\n^XA~TA000~JSN^LT0^MNW^MTD^PON^PMN^LH0,0^JMA^PR4,4~SD15^JUS^LRN^CI0^XZ\n^XA\n^MMT\n^PW448\n^LL0153\n^LS0\n^BY2,3,61^FT41,141^BCN,,N,N\n^FD>:" + barcode + "^FS\n^FT41,38^A0N,23,24^FH\\^FD" + text1 + "^FS\n^FT41,69^A0N,23,24^FH\\^FD" + text2 + "^FS\n^PQ" + qtyToPrint + ",0,1,Y^XZ";
        case 'item_warehouse_location_label':
            // Generate raw code to print two lines of large text and no barcode
            var text1 = arguments[2];
            var text2 = arguments[3];
            return "\u0010CT~~CD,~CC^~CT~\n^XA~TA000~JSN^LT0^MNW^MTD^PON^PMN^LH0,0^JMA^PR4,4~SD15^JUS^LRN^CI0^XZ\n^XA\n^MMT\n^PW448\n^LL0153\n^LS0\n^FT20,71^A0N,45,45^FH\\^FD" + text1 + "^FS\n^FT20,131^A0N,45,45^FH\\^FD" + text2 + "^FS\n^PQ" + qtyToPrint + ",0,1,Y^XZ";
    }

    return "";
}

// Removes new lines and ^ characters from the input value
function cleanZPLValue(value) {
    value = String(value);
    value = value.replace(/\^/g, '');
    value = value.replace(/\n/g, '');
    return value;
}


// http://www.gtin.info/
// 
// The GTIN is a globally unique 14-digit number used to identify trade items, products, or services. GTIN is also an umbrella term that refers to the entire family of UCC.EAN data structures. The entire family of data structures within the GTIN is:
// 
// GTIN-8 (EAN-8)
// GTIN-12 (UPC)
// GTIN-13 (EAN-13)
// GTIN-14 (EAN/UCC-128 or ITF-14)
// 
// Returns the type of barcode represented, or false if it is not a valid barcode (ie does not have the correct checkdigit at the end)
// Possible return values are (false, 'EAN-8', 'UPC-A', 'EAN-13', 'ITF-14')
function getBarcodeType(value) {
    value = String(value);

    // Make sure value is all digits and between 8 and 14 characters long
    var regex = new RegExp('^[0-9]{8,14}$');
    if (!regex.test(value)) {
        return false;
    }
    var orig_length = value.length;

    // Left pad value with '0's so it is 14 characters long
    while (value.length < 14) {
        value = '0' + value;
    }

    // Go through the first 13 digits
    var sum = 0;
    var ch_array = value.split('');
    for (var i = 0; i < 13; i++) {
        if ((i % 2) == 0) {
            sum += (+ch_array[i]) * 3;
        } else {
            sum += (+ch_array[i]);
        }
    }

    var check_digit = (10 - (sum % 10)) % 10;

    if (check_digit != (+ch_array[13])) {
        return false;
    }

    if (orig_length == 8) {
        return 'EAN-8';
    } else if (orig_length == 12) {
        return 'UPC-A';
    } else if (orig_length == 13) {
        return 'EAN-13';
    } else if (orig_length == 14) {
        return 'ITF-14';
    } else {
        return false;
    }
}