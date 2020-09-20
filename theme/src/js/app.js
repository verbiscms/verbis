/**
 * app.js
 * All custom JS for theme stored here.
 * @author Ainsley Clark
 * @author URL:   https://reddico.co.uk
 * @author Email: ainsley@reddico.co.uk
 */

/**
 * Require * Import
 * 
 */

//Local
import {$, getSiblings} from './scripts/helpers';
require('./scripts/polyfills');
require('./scripts/three')
require('./scripts/forms')
require('./scripts/modal')

//Vendor
import LazyLoad from 'vanilla-lazyload';
import Multiple from 'multiple.js'

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

// Lazyload Instance
let lazyLoadInstance = new LazyLoad({
	elements_selector: '.lazy'
	// ... more custom settings?
});

// Remove lazy display none
const lazyItems = document.querySelectorAll(".lazy")
lazyItems.forEach(item => {
	item.classList.add("lazy-show");
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

const scrollLinks = document.querySelectorAll('a[href^="/#"]');
scrollLinks.forEach(anchor => {
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
		hamburger.classList.add('hamburger-scrolled');
	} else {
		header.classList.remove('header-scrolled');
		nav.classList.remove('nav-scrolled');
		hamburger.classList.remove('hamburger-scrolled');
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

/*
 * Three
 * Removes classes once a link is clicked.
 */


/*
 * Share
 * Adds class when card share button is hovered.
 */

const shareBtns = document.querySelectorAll(".share-btn");
shareBtns.forEach(btn => {
	btn.addEventListener("mouseover", e => {
		e.preventDefault();

		const share = btn.closest(".share")
		let isOpen = share.getAttribute("data-open") === "true";

		if (!isOpen) {
			share.classList.add("share-active")
			share.setAttribute("data-open", "true")
		} else {
			share.classList.remove("share-active")
			share.setAttribute("data-open", "false")
		}
	});
});

// Link handler for when share items are clicked
const shareItems = document.querySelectorAll(".share-item:not(.share-btn)");
shareItems.forEach(btn => {
	btn.addEventListener("click", e => {
		e.preventDefault();

		console.log("clicked");

		const url = btn.getAttribute("data-link"),
			link = document.createElement('a');

		console.log(url);

		link.href = url;
		link.target = "_blank"
		document.body.appendChild(link);
		link.click();
		document.body.removeChild(link);
	});
});



