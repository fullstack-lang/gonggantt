// generated by ng_file_service_ts
import { Injectable, Component, Inject } from '@angular/core';
import { HttpClientModule } from '@angular/common/http';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { DOCUMENT, Location } from '@angular/common'

/*
 * Behavior subject
 */
import { BehaviorSubject } from 'rxjs';
import { Observable, of } from 'rxjs';
import { catchError, map, tap } from 'rxjs/operators';

import { BarDB } from './bar-db';

// insertion point for imports
import { LaneDB } from './lane-db'

@Injectable({
  providedIn: 'root'
})
export class BarService {

  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  BarServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private barsUrl: string

  constructor(
    private http: HttpClient,
    private location: Location,
    @Inject(DOCUMENT) private document: Document
  ) {
    // path to the service share the same origin with the path to the document
    // get the origin in the URL to the document
    let origin = this.document.location.origin

    // if debugging with ng, replace 4200 with 8080
    origin = origin.replace("4200", "8080")

    // compute path to the service
    this.barsUrl = origin + '/api/github.com/fullstack-lang/gonggantt/go/v1/bars';
  }

  /** GET bars from the server */
  getBars(): Observable<BarDB[]> {
    return this.http.get<BarDB[]>(this.barsUrl)
      .pipe(
        tap(_ => this.log('fetched bars')),
        catchError(this.handleError<BarDB[]>('getBars', []))
      );
  }

  /** GET bar by id. Will 404 if id not found */
  getBar(id: number): Observable<BarDB> {
    const url = `${this.barsUrl}/${id}`;
    return this.http.get<BarDB>(url).pipe(
      tap(_ => this.log(`fetched bar id=${id}`)),
      catchError(this.handleError<BarDB>(`getBar id=${id}`))
    );
  }

  //////// Save methods //////////

  /** POST: add a new bar to the server */
  postBar(bardb: BarDB): Observable<BarDB> {

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)
    let _Lane_Bars_reverse = bardb.Lane_Bars_reverse
    bardb.Lane_Bars_reverse = new LaneDB

    return this.http.post<BarDB>(this.barsUrl, bardb, this.httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        bardb.Lane_Bars_reverse = _Lane_Bars_reverse
        this.log(`posted bardb id=${bardb.ID}`)
      }),
      catchError(this.handleError<BarDB>('postBar'))
    );
  }

  /** DELETE: delete the bardb from the server */
  deleteBar(bardb: BarDB | number): Observable<BarDB> {
    const id = typeof bardb === 'number' ? bardb : bardb.ID;
    const url = `${this.barsUrl}/${id}`;

    return this.http.delete<BarDB>(url, this.httpOptions).pipe(
      tap(_ => this.log(`deleted bardb id=${id}`)),
      catchError(this.handleError<BarDB>('deleteBar'))
    );
  }

  /** PUT: update the bardb on the server */
  updateBar(bardb: BarDB): Observable<BarDB> {
    const id = typeof bardb === 'number' ? bardb : bardb.ID;
    const url = `${this.barsUrl}/${id}`;

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)
    let _Lane_Bars_reverse = bardb.Lane_Bars_reverse
    bardb.Lane_Bars_reverse = new LaneDB

    return this.http.put<BarDB>(url, bardb, this.httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        bardb.Lane_Bars_reverse = _Lane_Bars_reverse
        this.log(`updated bardb id=${bardb.ID}`)
      }),
      catchError(this.handleError<BarDB>('updateBar'))
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