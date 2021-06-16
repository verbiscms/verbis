/**
 * app.js
 * All custom JS for theme stored here.
 * @author Ainsley Clark
 * @author URL:   https://www.ainsleyclark.com
 * @author Email: info@ainsleyclark.com
 */

/**
 * Require * Import
 * 
 */

//Local
import {$, getSiblings} from './scripts/helpers';
require('./scripts/polyfills');

//Vendor
import hljs from 'highlight.js';
import 'highlight.js/styles/atelier-forest-light.css';
import handlebars from 'highlight.js/lib/languages/handlebars';
import go from 'highlight.js/lib/languages/go';
import assembly from 'highlight.js/lib/languages/x86asm';

/**
 * Variables
 *
 */
const html = document.querySelector('html'),
	header = document.querySelector('.header'),
	nav = document.querySelector('.nav'),
	hamburger = document.querySelector('.hamburger');

/*
 * Remove No JS Body Class
 *
 */
html.classList.remove('no-js');
html.classList.add('js');

const brPlugin = {
	"before:highlightBlock": ({ block }) => {
		block.innerHTML = block.innerHTML.replace(/\n/g, '').replace(/<br[ /]*>/g, '\n');
	},
	"after:highlightBlock": ({ result }) => {
		result.value = result.value.replace(/\n/g, "<br>");
	}
};

// how to use it
hljs.addPlugin(brPlugin);
/*
 * Highlight JS
 *
 */
hljs.registerLanguage('handlebars', handlebars);
hljs.registerLanguage('go', go);
hljs.registerLanguage('assembly', assembly);
hljs.highlightAll();

/*
 * Scroll To Anchor
 * Targets all links with # anchor & adds smooth scrolling
 *
 */
if (header) {
	let headerOffset = header.offsetHeight;

	window.addEventListener('resize', function(){
		headerOffset = header.offsetHeight;
	});

	document.querySelectorAll('a[href^="#"]').forEach(anchor => {
		anchor.addEventListener('click', function (e) {
			e.preventDefault();

			let offset = headerOffset,
				section = $(anchor.getAttribute('href')),
				elementPosition = section.offsetTop,
				offsetPosition = elementPosition - offset;

			window.scrollTo({
				top: offsetPosition,
				behavior: 'smooth'
			});
		});
	});
}

/*
 * Scroll
 * Adds header & nav classes after a certain scroll amount determined by scrollPos.
 *
 */
// const scrollPos = 100;
// if (header && nav) {
// 	window.addEventListener('scroll', function() {
// 		if (window.pageYOffset > scrollPos) {
// 			header.classList.add('header-scrolled');
// 			nav.classList.add('nav-scrolled');
// 		} else {
// 			header.classList.remove('header-scrolled');
// 			nav.classList.remove('nav-scrolled');
// 		}
// 	});
// }

/*
 * Nav Click
 * Removes classes once a link is clicked.
 *
 */

// Remove active classes when clicked.
const links = document.querySelectorAll('.header .nav .nav-item a');
links.forEach(link => {
	link.addEventListener('click', e => {
		if (window.innerWidth < 1025) {
			header.classList.remove('header-active');
			nav.classList.remove('nav-mobile-active');
			$('#hamburger-check').checked = '';
		}
	});
});

/*
 * Tabs
 * Handler for click of tabs, show and hide.
 *
 */
const tabs = document.querySelectorAll(".tab");
tabs.forEach(tab => {
	tab.addEventListener("click", e => {
		e.preventDefault();
		const attr = tab.getAttribute("data-tab"),
			panel = document.querySelector(attr);

		if (!panel) {
			console.warn("No panel exists with the attribute: " + attr);
			return;
		}

		document.querySelector(".tab-active").classList.remove("tab-active");
		tab.classList.add("tab-active");
		document.querySelector(".tab-panel-active").classList.remove("tab-panel-active");
		panel.classList.add("tab-panel-active");
	});
});

/*
 * Stack Frames
 * Handler for click of stack frame, change code panel.
 *
 */
const frames = document.querySelectorAll(".stack-frame-group");
frames.forEach(frame => {
	frame.addEventListener("click", e => {
		e.preventDefault();
		const attr = frame.getAttribute("data-stack"),
			stack = document.querySelector(attr);

		if (!stack) {
			console.warn("No stack exists with the attribute: " + attr);
			return;
		}

		document.querySelector(".stack-viewer-active").classList.remove("stack-viewer-active");
		stack.classList.add("stack-viewer-active");
		document.querySelector(".stack-frame-group-active").classList.remove("stack-frame-group-active");
		frame.classList.add("stack-frame-group-active");
	});
});

/*
 * Expand Vendor Frames
 * Handler for click of stack frame, change code panel.
 *
 */
const showVendorBtns = document.querySelectorAll(".stack-vendor-show");
showVendorBtns.forEach(btn => {
	btn.addEventListener("click", e => {
		e.preventDefault();

		const vendorFrames = document.querySelectorAll(".stack-frame-group-hidden");
		vendorFrames.forEach(frame => {
			frame.classList.remove("stack-frame-group-hidden")
		});
	})
});