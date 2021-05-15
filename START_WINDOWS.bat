rem ARG1=Telegram_BotCode  ARG2=GroupChatID ARG3=DistrictCode ARG4=AgeCheck
rem BotCode is unique API Key
rem Group chat id is @NameOfGroup
rem District Code is Cowin District code. (One way to find to search in Cowin website and then capture HTTP request of Cowin)
rem Example: Disctrict code of Gujarat-Kutch = 170, Maharashtra-Pune = 363
rem ArgeCheck is use for Getting update based on age. 18 = Get update of 18+, 45= get update of 45+ vaccine slots

rem covid19-vaccine-tracker.exe TelegramBotCode GroupChatID DistrictCode AgeCheck

covid19-vaccine-tracker.exe arg1 arg2 arg3 arg4
