import speech_recognition as sr
recognizer = sr.Recognizer()
with sr.Microphone() as source:
    audio_data = recognizer.listen(source)
    text = recognizer.recognize_google(audio_data)
