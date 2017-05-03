import {Component} from '@angular/core';

declare const Plaid: any;

@Component({
  selector: 'fountain-app',
  template: require('./hello.html')
})
export class HelloComponent {
  constructor() {
    
  var linkHandler = Plaid.create({
       env: 'sandbox',
       clientName: 'Plaid Sandbox',
       // Replace '<PUBLIC_KEY>' with your own `public_key`
       key: 'd75c9186402d2aff8219dab858c847',
       product: ['auth'],
       onSuccess: function(public_token, metadata) {
         // Send the public_token to your app server here.
         // The metadata object contains info about the
         // institution the user selected and the
         // account_id, if selectAccount is enabled.
         
    

        var xhr = new XMLHttpRequest();
        xhr.open("POST", '/login?token=' + public_token, true);
        xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8");
        xhr.send();

        
       },
       onExit: function(err, metadata) {
         // The user exited the Link flow.
         if (err != null) {
        console.log("Error: " + err);

         }
         // metadata contains information about the
         // institution that the user selected and the
         // most recent API request IDs. Storing this
         // information can be helpful for support.
       }
      });
      // Trigger the standard institution select view
       linkHandler.open();
  }
}
