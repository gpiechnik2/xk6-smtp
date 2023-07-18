import smtp from 'k6/x/smtp';


export default function () {
    const mailOptions = {
        subject: "Test subject",
        message: "Test message",
        udw: ["udwRecipient@gmail.com"]
    }

    smtp.sendMail(
        "smtp.gmail.com", 
        "587", 
        "sender@gmail.com", 
        "senderPassword", 
        "recipient@gmail.com",
        mailOptions
    )
}
