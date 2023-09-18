"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.addMonths = exports.XRPL_NETWORK = void 0;
exports.XRPL_NETWORK = "s.altnet.rippletest.net:51233";
const addMonths = (date, months) => {
    date.setMonth(date.getMonth() + months);
    return date;
};
exports.addMonths = addMonths;
