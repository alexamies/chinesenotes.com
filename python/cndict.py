# -*- coding: utf-8 -*-
"""
Utility loads the NTI Reader dictionary into a Python dictionary.

Simplified and traditional words are keys.
"""
import codecs


DICT_FILE_NAME = '../data/words.txt'
ACCENT_MAP = {u"ā": u"a", u"á": u"a", u"ǎ": u"a", u"à": u"a",
             u"ē": u"e", u"é": u"e", u"ě": u"e", u"è": u"e",
             u"ī": u"i", u"í": u"i", u"ǐ": u"i", u"ì": u"i",
             u"ō": u"o", u"ó": u"o", u"ǒ": u"o", u"ò": u"o",
             u"ū": u"u", u"ú": u"u", u"ǔ": u"u", u"ù": u"u",
             u"ǖ": u"u", u"ǘ": u"u", u"ǚ": u"u", u"ǜ": u"u"
            }


def OpenDictionary():
  """Reads the dictionary into memory
  """
  print "Opening the Chinese Notes dictionary"
  wdict = {}
  with codecs.open(DICT_FILE_NAME, 'r', "utf-8") as f:
    for line in f:
      line = line.strip()
      if not line:
        continue
      fields = line.split('\t')
      if fields and len(fields) >= 10:
        entry = {}
        entry['id'] = fields[0]
        entry['simplified'] = fields[1]
        entry['traditional'] = fields[2]
        entry['pinyin'] = fields[3]
        entry['english'] = fields[4]
        entry['grammar'] = fields[5]
        if fields and len(fields) >= 15 and fields[14] != '\\N':
          entry['notes'] = fields[14]
        traditional = entry['traditional']
        key = entry['simplified']
        if key not in wdict:
          entry['other_entries'] = []
          wdict[key] = entry
          if traditional != '\\N':
          	wdict[traditional] = entry
        else:
          wdict[key]['other_entries'].append(entry)
          if traditional != '\\N':
            if traditional in wdict:
              wdict[traditional]['other_entries'].append(entry)
            else:
              entry['other_entries'] = []
              wdict[traditional] = entry
  print "OpenDictionary completed with %d entries" % len(wdict)
  return wdict


def PersonName(wdict, trad):
  """Converts the traditional Chinese string into a person's name in English
  """
  english = []
  for i in range(len(trad)):
    t = trad[i]
    if t in wdict:
      entry = wdict[t]
      #print "%d, %s" % (i, t)
      if "pinyin" in entry:
        english.append(entry["pinyin"])
        #print "%d, len english %d" % (i, len(english))
        for key in ACCENT_MAP:
          english[i] = english[i].replace(key, ACCENT_MAP[key])
      else:
        english.append(t)
    else:
      #print "%s not in dictionary" % t
      english.append(t)

  english[0] = english[0].title()
  if len(english) > 1:
    english[1] = english[1].title()
  engStr = english[0] + " "
  for j in range (1, len(english)):
    engStr += english[j]
  return engStr


def main():
  wdict = OpenDictionary()
  PersonName(wdict, u"李穀")


if __name__ == "__main__":
  main()
