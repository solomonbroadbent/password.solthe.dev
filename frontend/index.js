import { PasswordGeneratorClient } from './generated/password-generator_grpc_web_pb.js';
import { PasswordRequest } from './generated/password-generator_grpc_web_pb.js'; 

// Now you can use PasswordGeneratorClient and PasswordRequest in your code
const client = new proto.password_generator.PasswordGeneratorClient('http://password.solthe.dev');
const request = new proto.password_generator.PasswordRequest();

document.getElementById('passwordGeneratorButton').addEventListener('click', () => {
	const request = new proto.password_generator.PasswordRequest();

  // Example usage: client.generatePassword(request, metadata, callback);
  client.generatePassword(request, null, (err, response) => {
    if (!err) {
      const generatedPassword = response.getPassword();
      alert('Generated password: ' + generatedPassword);
    } else {
      console.error('Error:', err.message);
    }
  });
});

// Example usage: client.generatePassword(request, metadata, callback);
client.generatePassword(request, null, (err, response) => {
  if (!err) {
    console.log('Generated password:', response.getPassword());
  } else {
    console.error('Error:', err.message);
  }
});
