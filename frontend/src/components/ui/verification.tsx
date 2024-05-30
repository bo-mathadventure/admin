"use server";

import nodemailer from 'nodemailer';


export const sendEmail = async (email: string) => {
   
    const { SMTP_EMAIL, SMTP_PASSWORD } = process.env;

    console.log(SMTP_EMAIL, SMTP_PASSWORD);

    const trasnport = nodemailer.createTransport({
        service: 'gmail',
        auth: {
            user: SMTP_EMAIL,
            pass: SMTP_PASSWORD
        }
    });
    try {
        const result = await trasnport.verify();
        console.log('Connected to email server', result);
    }
    catch (error) {
        console.error('Failed to connect to email server', error);
    }
    try {
        const info = await trasnport.sendMail({
            from: SMTP_EMAIL,
            to: email,
            subject: "Bestätigungsmail",
            text: "Bitte bestätigen Sie Ihre E-Mail-Adresse  Klicken Sie auf den folgenden Link, um Ihre E-Mail-Adresse zu bestätigen: " +"http://localhost:5656/verifyaccount/"+email,
        });
        console.log('Email sent', info);
    }
    catch (error) {
        console.error('Failed to send email', error);
    }
}

