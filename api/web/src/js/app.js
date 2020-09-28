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
import 'highlight.js/styles/atom-one-dark.css';
import handlebars from 'highlight.js/lib/languages/handlebars';
import go from 'highlight.js/lib/languages/go';


/**
 * Variables
 *
 */
const html = $('html'),
	header = $('.header'),
	nav = $('.nav'),
	hamburger = $('.hamburger');

/*
 * Highlight JS
 *
 */
hljs.registerLanguage('handlebars', handlebars);
hljs.registerLanguage('go', go);
hljs.initHighlightingOnLoad();

/*
 * Remove No JS Body Class
 *
 */
html.classList.remove('no-js');
html.classList.add('js');

/*
 * Scroll To Anchor
 * Targets all links with # anchor & adds smooth scrolling
 *
 */
let headerOffset = header.offsetHeight; 

window.addEventListener('resize', function(){
	headerOffset = header.offsetHeight;
});

$('a[href^="#"]').forEach(anchor => {
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

/*
 * Scroll
 * Adds header & nav classes after a certain scroll amount determined by scrollPos.
 *
 */
const scrollPos = 100;
window.addEventListener('scroll', function() {
	if (window.pageYOffset > scrollPos) {
		header.classList.add('header-scrolled');
		nav.classList.add('nav-scrolled');
	} else {
		header.classList.remove('header-scrolled');
		nav.classList.remove('nav-scrolled');
	}
});

/*
 * Nav Click
 * Removes classes once a link is clicked.
 */

// Remove active classes when clicked.
const links = $('.header .nav .nav-item a');
links.forEach(link => {
	link.addEventListener('click', e => {
		if (window.innerWidth < 1025) {
			header.classList.remove('header-active');
			nav.classList.remove('nav-mobile-active');
			$('#hamburger-check').checked = '';
		}
	});
});

