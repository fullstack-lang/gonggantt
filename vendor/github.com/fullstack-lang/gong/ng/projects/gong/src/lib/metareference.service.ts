// generated by ng_file_service_ts
import { Injectable, Component, Inject } from '@angular/core';
import { HttpClientModule, HttpParams } from '@angular/common/http';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { DOCUMENT, Location } from '@angular/common'

/*
 * Behavior subject
 */
import { BehaviorSubject } from 'rxjs';
import { Observable, of } from 'rxjs';
import { catchError, map, tap } from 'rxjs/operators';

import { MetaReferenceDB } from './metareference-db';

// insertion point for imports
import { MetaDB } from './meta-db'

@Injectable({
  providedIn: 'root'
})
export class MetaReferenceService {

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  MetaReferenceServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private metareferencesUrl: string

  constructor(
    private http: HttpClient,
    @Inject(DOCUMENT) private document: Document
  ) {
    // path to the service share the same origin with the path to the document
    // get the origin in the URL to the document
    let origin = this.document.location.origin

    // if debugging with ng, replace 4200 with 8080
    origin = origin.replace("4200", "8080")

    // compute path to the service
    this.metareferencesUrl = origin + '/api/github.com/fullstack-lang/gong/go/v1/metareferences';
  }

  /** GET metareferences from the server */
  getMetaReferences(GONG__StackPath: string = ""): Observable<MetaReferenceDB[]> {

	let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    return this.http.get<MetaReferenceDB[]>(this.metareferencesUrl, { params: params })
      .pipe(
        tap(_ => this.log('fetched metareferences')),
        catchError(this.handleError<MetaReferenceDB[]>('getMetaReferences', []))
      );
  }

  /** GET metareference by id. Will 404 if id not found */
  getMetaReference(id: number): Observable<MetaReferenceDB> {
    const url = `${this.metareferencesUrl}/${id}`;
    return this.http.get<MetaReferenceDB>(url).pipe(
      tap(_ => this.log(`fetched metareference id=${id}`)),
      catchError(this.handleError<MetaReferenceDB>(`getMetaReference id=${id}`))
    );
  }

  /** POST: add a new metareference to the server */
  postMetaReference(metareferencedb: MetaReferenceDB, GONG__StackPath: string): Observable<MetaReferenceDB> {

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)
    let _Meta_MetaReferences_reverse = metareferencedb.Meta_MetaReferences_reverse
    metareferencedb.Meta_MetaReferences_reverse = new MetaDB

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

	return this.http.post<MetaReferenceDB>(this.metareferencesUrl, metareferencedb, httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        metareferencedb.Meta_MetaReferences_reverse = _Meta_MetaReferences_reverse
        this.log(`posted metareferencedb id=${metareferencedb.ID}`)
      }),
      catchError(this.handleError<MetaReferenceDB>('postMetaReference'))
    );
  }

  /** DELETE: delete the metareferencedb from the server */
  deleteMetaReference(metareferencedb: MetaReferenceDB | number, GONG__StackPath: string): Observable<MetaReferenceDB> {
    const id = typeof metareferencedb === 'number' ? metareferencedb : metareferencedb.ID;
    const url = `${this.metareferencesUrl}/${id}`;

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.delete<MetaReferenceDB>(url, httpOptions).pipe(
      tap(_ => this.log(`deleted metareferencedb id=${id}`)),
      catchError(this.handleError<MetaReferenceDB>('deleteMetaReference'))
    );
  }

  /** PUT: update the metareferencedb on the server */
  updateMetaReference(metareferencedb: MetaReferenceDB, GONG__StackPath: string): Observable<MetaReferenceDB> {
    const id = typeof metareferencedb === 'number' ? metareferencedb : metareferencedb.ID;
    const url = `${this.metareferencesUrl}/${id}`;

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)
    let _Meta_MetaReferences_reverse = metareferencedb.Meta_MetaReferences_reverse
    metareferencedb.Meta_MetaReferences_reverse = new MetaDB

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.put<MetaReferenceDB>(url, metareferencedb, httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        metareferencedb.Meta_MetaReferences_reverse = _Meta_MetaReferences_reverse
        this.log(`updated metareferencedb id=${metareferencedb.ID}`)
      }),
      catchError(this.handleError<MetaReferenceDB>('updateMetaReference'))
    );
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation', result?: T) {
    return (error: any): Observable<T> => {

      // TODO: send the error to remote logging infrastructure
      console.error(error); // log to console instead

      // TODO: better job of transforming error for user consumption
      this.log(`${operation} failed: ${error.message}`);

      // Let the app keep running by returning an empty result.
      return of(result as T);
    };
  }

  private log(message: string) {

  }
}