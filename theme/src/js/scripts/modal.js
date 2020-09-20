/**
 * modal.js
 * Pull in information from modals & display.
 * @author Ainsley Clark
 * @author URL:   https://reddico.co.uk
 * @author Email: ainsley@reddico.co.uk
 */

// Opens modal
const modalBtns = document.querySelectorAll(".modal-btn");
modalBtns.forEach(function(button) {
    const id = button.getAttribute('data-modal'),
        modal = document.querySelector(id);

    button.addEventListener('click', function() {
        if (id === '#termsModal') {
            addTermsText(button.getAttribute('data-html'), button.innerHTML);
        }
        modal.classList.add('modal-open');
    });
});

// Closes modal
const modalHide = document.querySelectorAll(".modal-hide");
modalHide.forEach(function(el) {
    el.addEventListener('click', function() {
        el.closest('.modal').classList.remove('modal-open');
    });
});

//Pulls in terms & conditions text from html files.
function addTermsText(htmlFile, title) {

    let request = new XMLHttpRequest(),
        filePath = '/assets/modals/' + htmlFile,
        modalTitle = document.querySelector('#termsModal .terms-heading'),
        modalBody = document.querySelector('#termsModal .modal-body');

    request.open('GET', filePath, true);
    request.onload = function() {
        if (request.status >= 200 && request.status < 400) {
            const resp = request.responseText;
            modalTitle.innerHTML = title;
            modalBody.innerHTML = resp;
        }
    };
    request.onerror = function() {
        modalBody.innerHTML = 'Sorry something went wrong, please try again.';
        const response = JSON.parse(this.response);
        console.log(response);
    };
    request.send();
}