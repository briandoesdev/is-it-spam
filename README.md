# Is It Spam?

*This is a submission for [Twilio Challenge v24.06.12](https://dev.to/challenges/twilio)*

## What I Built
Everyone who has dealt with spam calls knows how stressful it can be getting a call from an unknown number. Maybe you are waiting for a call from a job application, or your childs teacher, or the doctor. There are a thousand reasons you may need to keep yourself available to unknown callers. But you don't want to pickup spam callers, letting them know your number is active and that you answer, only to start getting more spam calls.

Well what I've built is a product that utilizes Twilio and AI to help you determine if a caller is potentially spam or not. By providing the callers number, "Is It Spam?" leverages Twilio's Lookup API and add-on marketplace to retrieve data about the caller. That data is then parsed by GPT-3.5 Turbo to determine if the caller is spam, using attributes associated to the number like line type, name, historical data, etc... This is then presented back to the user either via the web interface.

**This service will only be public for the duration of the contest and judging. Due to expenses with using the Twilio and OpenAI API's, I cannot afford to keep this running. The source code is available, all you need to run this is the latest version of Go (go1.22.4), a Twilio account, and an OpenAI API Key.**

## Demo
Demo Link: [Is It Spam?](https://is-it-spam.brians.land)
GitHub Link: [briandoesdev/is-it-spam](https://github.com/briandoesdev)
<!-- Screenshots -->

![Lookup Page](https://dev-to-uploads.s3.amazonaws.com/uploads/articles/jryw74yjxyhflhwwz969.png)

![Lookup Response](https://dev-to-uploads.s3.amazonaws.com/uploads/articles/yr9v3fdmi476xfyy4lwk.png)

## Twilio and AI
I use the Twilio Lookup API, along with a marketplace add-on, to look up data about the provided number. This is then passed to the OpenAI API to parse and summarize the likelyhood this is spam via specified attributes.

## Additional Prize Categories
- Impactful Innovators: Whilst this could be subjective, I'd say helping prevent spam calls, and knowing who is calling from an unknown number can provide a societal impact. I know that I've gotten less stress from calls since developing this service.

<!-- Don't forget to add a cover image (if you want). -->

<!-- Thanks for participating! -->

### Credits
HTMX Loading Indicator: [u/M8Ir88outOf8](https://www.reddit.com/r/htmx/comments/1blwnc4/tip_of_the_day_unobtrusive_global_loading/)

