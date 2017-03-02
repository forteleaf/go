#!/usr/bin/env python
import sys,re,operator,string

stack = []
heap = []

def read_file():
   f = open(stack.pop())
   stack.append([f.read()])
   f.close()

def filter_chars():
   #정규표현식으로 처리
   stack.append(re.compile('[\W_]+'))
   #그결과를 스택에 넣는다
   stack.append([stack.pop().sub(' ',stack,pop()[0]).lower()])

def scan():
   stack.extend(stack.pop()[0].split())

def remove_stop_words():
   f = open('../stop_words.txt')
   stack.append(f.read().split(','))
   f.close()
   stack[-1].extend(lsit(string.ascii_lowercase))
   heap['stop_words'] =stack.pop()
   heap['words'] = []
   while len(stack) > 0:
      if stack[-1] in heap['stop_words']:
         stack.pop()
      else:
         heap['words'].append(stack.pop())
   stack.extend(heap['words'])
   del heap['stop_words']; del heap['words']

def frequencies():
   heap['word_freqs'] = {}
   while len(stack) >0:
      if stack[-1] in heap['word_freqs']:
         stack.append(heap['word_freqs'][stack[-1]])
         stack.append(1)
         stack.append(stack.pop() + stack.pop())
      else:
         stack.append(1)
      heap['word_freqs'][stack.pop()] = stack.pop()
   
   #그 결과를 스텍에 넣는다.
   stack.append(heap['word_freqs'])
   del heap['word_freqs'] # 이변수는 더이상 필요하지 않다.

def sort():
   stack.extend(sorted(stack.pop().iteritems(),key=operator.itemgetter(1)))

stack.append(sys.argv[1])
read_file();filter_chars();scan();remove_stop_words()
frequencies(); sort()

stack.append(0)

while stack[-1] < 25 and len(stack)>1:
   heap['i'] = stack.pop()
   (w,f) = stack.pop();print(w,' - ', f)
   stack.append(heap['i']);stack.append(1)
   stack.append(stack.pop() + stack.pop())