"use server";

import nodemailer from 'nodemailer';


export const sendEmail = async (email: string, token : string ) => {
   
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
            text: email+ " has a token: " + token,
        });
        console.log('Email sent', info);
    }
    catch (error) {
        console.error('Failed to send email', error);
    }
}

