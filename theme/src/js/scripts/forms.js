/**
 * forms.js
 * Custom JS for form sends stored here.
 * @author Ainsley Clark
 * @author URL:   https://reddico.co.uk
 * @author Email: ainsley@reddico.co.uk
 */

/*
 * Send Form
 * Logic to send subscriber to backend.
 */

const sendBtns = document.querySelectorAll(".form-send");
sendBtns.forEach(btn => {
   btn.addEventListener("click", e => {
       e.preventDefault();

       const form = btn.closest(".form"),
           inputs = form.querySelectorAll('.form-input'),
           sendData = {};

       btn.classList.add("btn-loading");

       inputs.forEach(input => {
           const name = input.getAttribute('name');
           input.type === 'checkbox' ? sendData[name] = (input.checked ? 'true' : 'false') : sendData[name] = input.value;
       });

       // Send Email
       let request = new XMLHttpRequest();
       request.open('POST', '/ajax/subscribe', true);
       request.setRequestHeader('Content-Type', 'application/json');
       request.onload = function() {
           if (this.status >= 200 && this.status < 400) {
               form.classList.add("form-success");
           } else {

               validateForm(JSON.parse(this.response).data, form);
               setTimeout(e => {
                   btn.classList.remove('btn-loading');
               }, 300);
           }
       };
       request.onerror = function() {
           setTimeout(e => {
               btn.classList.remove('btn-loading');
           }, 300);
       };
       request.send(JSON.stringify(sendData));
   });
});

/*
 * Validate Form
 * Custom validation for fields.
 */
const validateForm = (data, form) => {
   data.forEach(datum => {
       const formInputConts = form.querySelector(".form-input-cont"),
           input = form.querySelector(`[name=${datum.key}]`),
           inputCont = input.closest(".form-input-cont"),
           message = inputCont.querySelector(".form-message");

       inputCont.classList.add("form-has-error");
       message.innerHTML = datum.message;
   });
}