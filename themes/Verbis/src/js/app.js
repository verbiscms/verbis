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
import LazyLoad from 'vanilla-lazyload';

/**
 * Variables
 *
 */
const html = $('html'),
	header = $('.header'),
	nav = $('.nav'),
	hamburger = $('.hamburger');

/*
 * Remove No JS Body Class
 *
 */
html.classList.remove('no-js');
html.classList.add('js');

/**
 * Vanilla Lazyload
 * 
 */
let lazyLoadInstance = new LazyLoad({
	elements_selector: '.lazy'
	// ... more custom settings?
});

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

