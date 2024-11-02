function sendMessage() {
    const input = document.getElementById('chat-input');
    const messagesContainer = document.getElementById('chat-messages');
    const messageText = input.value;
    if (messageText.trim() === '') {
        alert('Please enter a message.');
        return;
    }
    // Send message to the server
    fetch('/chat', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({text: messageText}),
    })
        .then(response => response.json())
        .then(data => {
            if (data.code === 0 && data.msg === 'ok') {
                // Display the response from the server
                const messageElement = document.createElement('div');
                messageElement.classList.add('message');
                messageElement.innerText = data.data; // Assuming the response contains the generated text in "data" field
                messagesContainer.appendChild(messageElement);
                // Scroll to the bottom of the messages container
                messagesContainer.scrollTop = messagesContainer.scrollHeight;
            } else {
                alert('Error: ' + data.msg);
            }
        })
        .catch(error => {
            console.error('Error:', error);
            alert('An error occurred while sending the message.');
        });
    // Clear the input field
    input.value = '';
}