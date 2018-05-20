# -*- coding: utf-8 -*-

import codecs

dname = '../corpus/'
text = """\n\n
Chinese text: This work was published before January 1, 1923, and is in the
public domain worldwide because the author died at least 100 years ago.
"""
num = [u'', u'一', u'二', u'三', u'四', u'五', u'六', u'七', u'八', u'九']
with codecs.open('temp.txt', 'w', "utf-8") as f:
  prefix = 'jinshu'
  for i in range(1,131):
    fname = ''
    j = i % 10
    k = i / 10
    if i < 10:
      scroll = u'%s' % (num[j])
    elif i < 20:
      scroll = u'十%s' % (num[j])
    elif i < 100:
      scroll = u'%s十%s' % (num[k], num[j])
    elif i < 110:
      scroll = u'一百〇%s' % (num[j])
    else:
      k = k % 10
      scroll = u'一百%s十%s' % (num[k], num[j])
    title = u'\t卷%s Volume %s: \n' % (scroll, i)
    # print('%d %d %s' % (j, k, title))
    if i < 10:
      fname = u'{0}/{1}00{2}.txt'.format(prefix, prefix, i)
      f.write(u'%s\t%s/%s00%d.html%s' % (fname, prefix, prefix, i, title))
    elif i < 100:
      fname = u'{0}/{1}0{2}.txt'.format(prefix, prefix, i)
      f.write(u'%s\t%s/%s0%d.html%s' % (fname, prefix, prefix, i, title))
    else:
      fname = u'{0}/{1}{2}.txt'.format(prefix, prefix, i)
      f.write(u'%s\t%s/%s%d.html%s' % (fname, prefix, prefix, i, title))
    path = '{0}{1}'.format(dname, fname)
    g = open(path, 'w')
    g.write(text)