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

import { MilestoneDB } from './milestone-db';

// insertion point for imports
import { GanttDB } from './gantt-db'

@Injectable({
  providedIn: 'root'
})
export class MilestoneService {

  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  MilestoneServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private milestonesUrl: string

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
    this.milestonesUrl = origin + '/api/github.com/fullstack-lang/gonggantt/go/v1/milestones';
  }

  /** GET milestones from the server */
  getMilestones(): Observable<MilestoneDB[]> {
    return this.http.get<MilestoneDB[]>(this.milestonesUrl)
      .pipe(
        tap(_ => this.log('fetched milestones')),
        catchError(this.handleError<MilestoneDB[]>('getMilestones', []))
      );
  }

  /** GET milestone by id. Will 404 if id not found */
  getMilestone(id: number): Observable<MilestoneDB> {
    const url = `${this.milestonesUrl}/${id}`;
    return this.http.get<MilestoneDB>(url).pipe(
      tap(_ => this.log(`fetched milestone id=${id}`)),
      catchError(this.handleError<MilestoneDB>(`getMilestone id=${id}`))
    );
  }

  //////// Save methods //////////

  /** POST: add a new milestone to the server */
  postMilestone(milestonedb: MilestoneDB): Observable<MilestoneDB> {

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)
    milestonedb.LanesToDisplayMilestoneUse = []
    let _Gantt_Milestones_reverse = milestonedb.Gantt_Milestones_reverse
    milestonedb.Gantt_Milestones_reverse = new GanttDB

    return this.http.post<MilestoneDB>(this.milestonesUrl, milestonedb, this.httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        milestonedb.Gantt_Milestones_reverse = _Gantt_Milestones_reverse
        this.log(`posted milestonedb id=${milestonedb.ID}`)
      }),
      catchError(this.handleError<MilestoneDB>('postMilestone'))
    );
  }

  /** DELETE: delete the milestonedb from the server */
  deleteMilestone(milestonedb: MilestoneDB | number): Observable<MilestoneDB> {
    const id = typeof milestonedb === 'number' ? milestonedb : milestonedb.ID;
    const url = `${this.milestonesUrl}/${id}`;

    return this.http.delete<MilestoneDB>(url, this.httpOptions).pipe(
      tap(_ => this.log(`deleted milestonedb id=${id}`)),
      catchError(this.handleError<MilestoneDB>('deleteMilestone'))
    );
  }

  /** PUT: update the milestonedb on the server */
  updateMilestone(milestonedb: MilestoneDB): Observable<MilestoneDB> {
    const id = typeof milestonedb === 'number' ? milestonedb : milestonedb.ID;
    const url = `${this.milestonesUrl}/${id}`;

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)
    milestonedb.LanesToDisplayMilestoneUse = []
    let _Gantt_Milestones_reverse = milestonedb.Gantt_Milestones_reverse
    milestonedb.Gantt_Milestones_reverse = new GanttDB

    return this.http.put<MilestoneDB>(url, milestonedb, this.httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        milestonedb.Gantt_Milestones_reverse = _Gantt_Milestones_reverse
        this.log(`updated milestonedb id=${milestonedb.ID}`)
      }),
      catchError(this.handleError<MilestoneDB>('updateMilestone'))
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
