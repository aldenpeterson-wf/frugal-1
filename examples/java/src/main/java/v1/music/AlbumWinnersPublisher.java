/**
 * Autogenerated by Frugal Compiler (1.19.2)
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *
 * @generated
 */

package v1.music;

import com.workiva.frugal.middleware.InvocationHandler;
import com.workiva.frugal.middleware.ServiceMiddleware;
import com.workiva.frugal.protocol.*;
import com.workiva.frugal.provider.FScopeProvider;
import com.workiva.frugal.transport.FScopeTransport;
import com.workiva.frugal.transport.FSubscription;
import org.apache.thrift.TException;
import org.apache.thrift.TApplicationException;
import org.apache.thrift.transport.TTransportException;
import org.apache.thrift.protocol.*;

import javax.annotation.Generated;
import java.util.logging.Logger;




/**
 * Scopes are a Frugal extension to the IDL for declaring PubSub
 * semantics. Subscribers to this scope will be notified if they win a contest.
 * Scopes must have a prefix.
 */
@Generated(value = "Autogenerated by Frugal Compiler (1.19.2)", date = "2016-10-5")
public class AlbumWinnersPublisher {

	private static final String DELIMITER = ".";

	private final InternalAlbumWinnersPublisher target;
	private final InternalAlbumWinnersPublisher proxy;

	public AlbumWinnersPublisher(FScopeProvider provider, ServiceMiddleware... middleware) {
		target = new InternalAlbumWinnersPublisher(provider);
		proxy = (InternalAlbumWinnersPublisher) InvocationHandler.composeMiddlewareClass(target, InternalAlbumWinnersPublisher.class, middleware);
	}

	public void open() throws TException {
		target.open();
	}

	public void close() throws TException {
		target.close();
	}

	public void publishWinner(FContext ctx, Album req) throws TException {
		proxy.publishWinner(ctx, req);
	}

	protected static class InternalAlbumWinnersPublisher {

		private FScopeProvider provider;
		private FScopeTransport transport;
		private FProtocol protocol;

		protected InternalAlbumWinnersPublisher() {
		}

		public InternalAlbumWinnersPublisher(FScopeProvider provider) {
			this.provider = provider;
		}

		public void open() throws TException {
			FScopeProvider.Client client = provider.build();
			transport = client.getTransport();
			protocol = client.getProtocol();
			transport.open();
		}

		public void close() throws TException {
			transport.close();
		}

		public void publishWinner(FContext ctx, Album req) throws TException {
			String op = "Winner";
			String prefix = "v1.music.";
			String topic = String.format("%sAlbumWinners%s%s", prefix, DELIMITER, op);
			transport.lockTopic(topic);
			try {
				protocol.writeRequestHeader(ctx);
				protocol.writeMessageBegin(new TMessage(op, TMessageType.CALL, 0));
				req.write(protocol);
				protocol.writeMessageEnd();
				transport.flush();
			} finally {
				transport.unlockTopic();
			}
		}
	}
}
