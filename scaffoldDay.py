import argparse
import datetime
import os
import requests
import shutil
import sys

def main():
  now = datetime.datetime.now()

  parser = argparse.ArgumentParser(description='Script to setup the advent of code problem.')
  parser.add_argument('day', type=int, help='the day of the problem you want setup', default=now.day, nargs='?')
  parser.add_argument('year', type=int, help='the year of the problem you want setup', default=now.year, nargs='?')
  args = parser.parse_args()
  
  year = args.year
  day = args.day
  if year >15 and year <99:
    year +=2000

  if year > now.year or year < 2015:
    sys.exit("Invalid year {}".format(year))
  if day > 25 or day < 0:
    sys.exit("Invalid day {}".format(day))
  if year == now.year and ((now.month < 12) or (now.day < day)):
    sys.exit("Cannot predict the future ({}-12-{})".format(year,day))


  basePath = os.path.dirname(os.path.realpath(__file__))
  todayPath = os.path.join(basePath,str(year),"day{:02d}".format(day))
  templatePath = os.path.join(basePath,"dayX")
  
  if not os.path.exists(todayPath):
    os.makedirs(todayPath)

  print(todayPath)

  for filename in os.listdir(templatePath):
    if filename == "input.txt":
      continue
    source = os.path.join(templatePath, filename)
    target = os.path.join(todayPath, filename)
    if not os.path.exists(os.path.join(todayPath,filename)):
      shutil.copyfile(src=source, dst=target)

  inputFile = os.path.join(todayPath,"input.txt")
  if not os.path.exists(inputFile):
    getInput(year,day,inputFile)

  challengeURL = "https://adventofcode.com/{}/day/{}".format(year,day)
  print ("Challenge URL is:\n\n\t{}\n".format(challengeURL), file=sys.stderr)

def makePath(path):
  os.path.exists(path)
  os.mkdir()

def getInput(year,day,savePath):
  inputURL = "https://adventofcode.com/{}/day/{}/input".format(year,day)

  with open("session.cookie",'r') as sessionFile:
    session = sessionFile.read().replace('\n', '')
  if not session:
    sys.exit("Session not specified")
  
  cookies = {
    "session": session,
  }
  
  attemptsRemaining = 5
  while attemptsRemaining > 0:
    attemptsRemaining-=1

    result = requests.get(inputURL,cookies=cookies,stream=True)
    if result.status_code == 200:
      with open(savePath, 'wb') as inputFile:
        for chunk in result:
            inputFile.write(chunk)
      return 
    print("Failed to fetch input",result,file=sys.stderr)


if __name__ == "__main__":
  main()
