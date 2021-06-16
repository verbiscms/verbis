/**
 * helpers.js
 * Vanilla JS helpers.
 * @author Ainsley Clark
 * @author URL:   https://www.ainsleyclark.com
 * @author Email: info@ainsleyclark.com
 */

/**
 * Query Selector
 * Usage: $('.classes'), $1('.classname')
 * 
 */

// Select a list/single of matching elements, context is optional
export const $ = (selector, context) => {
    'use strict';

    const el = (context || document).querySelectorAll(selector);

    if (!el || el.length == 0) {
        console.warn(`${selector} element not found in DOM.`);
        return el;
    } else if (el.length == 1) {
        return el[0];
    }

    return el;
}

/**
 * Get Siblings
 * 
 */
export const getSiblings = (el, filter) => {
    var siblings = [];
    el = el.parentNode.firstChild;
    do { if (!filter || filter(el)) siblings.push(el); } while (el = el.nextSibling);
    return siblings;
}

export default {$, getSiblings};