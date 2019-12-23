/*
 * Licensed  under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

/**
 *  @fileoverview  Interface definitions for objects read remotely.
 */

// Interface for JSON collection data loaded from AJAX call
export interface ICollection {
  GlossFile: string;
  Title: string;
}

// Interface for results loaded from AJAX call
export interface IDocSearchRestults {
  Collections: ICollection[];
  Documents: IDocument[];
  NumCollections: number;
  NumDocuments: number;
  Terms: ITerm[];
}

// Interface for JSON document data loaded from AJAX call
export interface IDocument {
  CollectionFile: string;
  CollectionTitle: string;
  GlossFile: string;
  MatchDetails: IMatchDetails;
  SimBigram: string;
  SimBitVector: string;
  SimTitle: string;
  Title: string;
}

// Interface for JSON dictionary entry data loaded from AJAX call
export interface IDictEntry {
  HeadwordId: number;
  Pinyin: string;
  Senses: IWordSense[];
}

// Interface for JSON match detailed data loaded from AJAX call
export interface IMatchDetails {
  ExactMatch: string;
  LongestMatch: string;
  Snippet: string;
}

// Interface for JSON query term data loaded from AJAX call
export interface ITerm {
  DictEntry: IDictEntry;
  QueryText: string;
  Senses: IWordSense[];
}

// Interface for JSON WordSense data loaded from AJAX call
export interface IWordSense {
  English: string;
  HeadwordId: string;
  Notes: string;
  Pinyin: string;
  Simplified: string;
  Traditional: string;
}
