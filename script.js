
// POPOVER LIBRARY
$('.popover-dismiss').popover({
    trigger: 'focus'
  })

  document.querySelector("button[type='submit']").addEventListener("click", function(event) {
    event.preventDefault(); // prevent the form from submitting
    var userInput = document.querySelector("textarea#user-input").value;
    console.log("Entered text: " + userInput);
});